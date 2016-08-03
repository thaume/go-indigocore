// Provides functionality to create an HTTP server from an adapter.
package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	. "github.com/stratumn/go/store/adapter"
	. "github.com/stratumn/go/store/segment"
	. "github.com/stratumn/go/store/segment/segmentvalidation"
)

const (
	DEFAULT_PORT    = ":5000"
	DEFAULT_VERBOSE = false
)

// A server configuration.
type Config struct {
	Port     string
	CertFile string
	KeyFile  string
	Verbose  bool
}

// A server.
type Server struct {
	config *Config
	router *httprouter.Router
}

// Implements HTTP server.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Makes a new server.
func New(adapter Adapter, config *Config) *Server {
	r := httprouter.New()
	c := Context{adapter, config}

	r.NotFound = NotFoundHandler{JSONHandler{&c, NotFound}}

	r.GET("/", JSONHandler{&c, Root}.ServeHTTP)
	r.POST("/segments", JSONHandler{&c, SaveSegment}.ServeHTTP)
	r.GET("/segments/:linkHash", JSONHandler{&c, GetSegment}.ServeHTTP)
	r.DELETE("/segments/:linkHash", JSONHandler{&c, DeleteSegment}.ServeHTTP)
	r.GET("/segments", JSONHandler{&c, FindSegments}.ServeHTTP)
	r.GET("/maps", JSONHandler{&c, GetMapIDs}.ServeHTTP)

	return &Server{config, r}
}

// Starts listening and serving.
func (s *Server) ListenAndServe() error {
	port := s.config.Port

	if port == "" {
		port = DEFAULT_PORT
	}

	log.Printf("listening on %s\n", port)

	if s.config.CertFile != "" && s.config.KeyFile != "" {
		return http.ListenAndServeTLS(port, s.config.CertFile, s.config.KeyFile, s)
	}

	return http.ListenAndServe(port, s)
}

// HTTP handle for 404s.
func NotFound(w http.ResponseWriter, r *http.Request, _ *Context, _ httprouter.Params) (interface{}, error) {
	return nil, &ErrNotFound
}

// HTTP handle for the root route.
func Root(w http.ResponseWriter, r *http.Request, c *Context, _ httprouter.Params) (interface{}, error) {
	info, err := c.Adapter.GetInfo()

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"adapter": info,
	}, nil
}

// HTTP handle to save a segment.
func SaveSegment(w http.ResponseWriter, r *http.Request, c *Context, p httprouter.Params) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)

	var segment Segment

	if err := decoder.Decode(&segment); err != nil {
		return nil, &ErrBadRequest
	}

	if err := Validate(&segment); err != nil {
		return nil, &ErrHTTP{err.Error(), 400}
	}

	if err := c.Adapter.SaveSegment(&segment); err != nil {
		return nil, err
	}

	return segment, nil
}

// HTTP handle to get a segment.
func GetSegment(w http.ResponseWriter, r *http.Request, c *Context, p httprouter.Params) (interface{}, error) {
	segment, err := c.Adapter.GetSegment(p.ByName("linkHash"))

	if err != nil {
		return nil, err
	}

	if segment == nil {
		return nil, &ErrNotFound
	}

	return segment, nil
}

// HTTP handle to delete a segment.
func DeleteSegment(w http.ResponseWriter, r *http.Request, c *Context, p httprouter.Params) (interface{}, error) {
	segment, err := c.Adapter.DeleteSegment(p.ByName("linkHash"))

	if err != nil {
		return nil, err
	}

	if segment == nil {
		return nil, &ErrNotFound
	}

	return segment, nil
}

// HTTP handle to show segments.
func FindSegments(w http.ResponseWriter, r *http.Request, c *Context, p httprouter.Params) (interface{}, error) {
	filter, errHTTP := parseFilter(r)

	if errHTTP != nil {
		return nil, errHTTP
	}

	segments, err := c.Adapter.FindSegments(filter)

	if err != nil {
		return nil, err
	}

	return segments, nil
}

// HTTP handle to show map ids.
func GetMapIDs(w http.ResponseWriter, r *http.Request, c *Context, p httprouter.Params) (interface{}, error) {
	pagination, errHTTP := parsePagination(r)

	if errHTTP != nil {
		return nil, errHTTP
	}

	mapIDs, err := c.Adapter.GetMapIDs(pagination)

	if err != nil {
		return nil, err
	}

	return mapIDs, nil
}

// Creates an adapter filter from a request.
func parseFilter(r *http.Request) (*Filter, error) {
	var tags []string

	pagination, errHTTP := parsePagination(r)

	if errHTTP != nil {
		return nil, errHTTP
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
