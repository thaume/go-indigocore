// Provides functionality to create an HTTP server from an adapter.
package httpserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/stratumn/go/jsonhttp"
	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/segment/segmentvalidation"
	. "github.com/stratumn/go/store/adapter"
)

const (
	DEFAULT_PORT    = ":5000"
	DEFAULT_VERBOSE = false
)

// A server context.
type context struct {
	adapter Adapter
	config  *jsonhttp.Config
}

// A server handle.
type handle func(http.ResponseWriter, *http.Request, httprouter.Params, *context) (interface{}, error)

// A server handler.
type handler struct {
	context *context
	handle  handle
}

func (h handler) serve(w http.ResponseWriter, r *http.Request, p httprouter.Params, _ *jsonhttp.Config) (interface{}, error) {
	return h.handle(w, r, p, h.context)
}

// Makes a new server.
func New(a Adapter, c *jsonhttp.Config) *jsonhttp.Server {
	server := jsonhttp.New(c)
	context := &context{a, c}

	server.Get("/", handler{context, root}.serve)
	server.Post("/segments", handler{context, saveSegment}.serve)
	server.Get("/segments/:linkHash", handler{context, getSegment}.serve)
	server.Delete("/segments/:linkHash", handler{context, deleteSegment}.serve)
	server.Get("/segments", handler{context, findSegments}.serve)
	server.Get("/maps", handler{context, getMapIDs}.serve)

	return server
}

// HTTP handle for the root route.
func root(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c *context) (interface{}, error) {
	info, err := c.adapter.GetInfo()

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"adapter": info,
	}, nil
}

// HTTP handle to save a segment.
func saveSegment(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c *context) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)

	var s Segment

	if err := decoder.Decode(&s); err != nil {
		return nil, &jsonhttp.ErrBadRequest
	}

	if err := Validate(&s); err != nil {
		return nil, &jsonhttp.ErrHTTP{err.Error(), 400}
	}

	if err := c.adapter.SaveSegment(&s); err != nil {
		return nil, err
	}

	return s, nil
}

// HTTP handle to get a segment.
func getSegment(w http.ResponseWriter, r *http.Request, p httprouter.Params, c *context) (interface{}, error) {
	s, err := c.adapter.GetSegment(p.ByName("linkHash"))

	if err != nil {
		return nil, err
	}

	if s == nil {
		return nil, &jsonhttp.ErrNotFound
	}

	return s, nil
}

// HTTP handle to delete a segment.
func deleteSegment(w http.ResponseWriter, r *http.Request, p httprouter.Params, c *context) (interface{}, error) {
	s, err := c.adapter.DeleteSegment(p.ByName("linkHash"))

	if err != nil {
		return nil, err
	}

	if s == nil {
		return nil, &jsonhttp.ErrNotFound
	}

	return s, nil
}

// HTTP handle to show segments.
func findSegments(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c *context) (interface{}, error) {
	filter, e := parseFilter(r)

	if e != nil {
		return nil, e
	}

	slice, err := c.adapter.FindSegments(filter)

	if err != nil {
		return nil, err
	}

	return slice, nil
}

// HTTP handle to show map ids.
func getMapIDs(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c *context) (interface{}, error) {
	pagination, e := parsePagination(r)

	if e != nil {
		return nil, e
	}

	slice, err := c.adapter.GetMapIDs(pagination)

	if err != nil {
		return nil, err
	}

	return slice, nil
}

// Creates an adapter filter from a request.
func parseFilter(r *http.Request) (*Filter, error) {
	var tags []string

	pagination, e := parsePagination(r)

	if e != nil {
		return nil, e
	}

	mapID := r.URL.Query().Get("mapId")
	prevLinkHash := r.URL.Query().Get("prevLinkHash")

	tagsStr := r.URL.Query().Get("tags")

	if tagsStr != "" {
		spacetags := strings.Split(tagsStr, " ")

		for _, t := range spacetags {
			tags = append(tags, strings.Split(t, "+")...)
		}
	}

	return &Filter{
		Pagination:   *pagination,
		MapID:        mapID,
		PrevLinkHash: prevLinkHash,
		Tags:         tags,
	}, nil
}

// Creates an adapter pagination from a request.
func parsePagination(r *http.Request) (*Pagination, error) {
	var err error

	offsetstr := r.URL.Query().Get("offset")
	offset := 0

	if offsetstr != "" {
		if offset, err = strconv.Atoi(offsetstr); err != nil || offset < 0 {
			return nil, &ErrOffset
		}
	}

	limitstr := r.URL.Query().Get("limit")
	limit := 0

	if limitstr != "" {
		if limit, err = strconv.Atoi(limitstr); err != nil || limit < 0 {
			return nil, &ErrLimit
		}
	}

	return &Pagination{
		Offset: offset,
		Limit:  limit,
	}, nil
}
