// Provides functionality to create an HTTP server from an
package httpserver

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	. "github.com/stratumn/go/fossilizer/adapter"
	"github.com/stratumn/go/jsonhttp"
)

const (
	DEFAULT_PORT               = ":6000"
	DEFAULT_NUM_RESULT_WORKERS = 8
	DEFAULT_MIN_DATA_LEN       = 32
	DEFAULT_MAX_DATA_LEN       = 64
	DEFAULT_VERBOSE            = false
)

// A server configuration.
type Config struct {
	jsonhttp.Config
	NumResultWorkers int
	MinDataLen       int
	MaxDataLen       int
}

// A server context.
type context struct {
	adapter Adapter
	config  *Config
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
func New(a Adapter, c *Config) *jsonhttp.Server {
	if c.NumResultWorkers < 1 {
		c.NumResultWorkers = DEFAULT_NUM_RESULT_WORKERS
	}

	s := jsonhttp.New(&c.Config)
	ctx := &context{a, c}

	s.Get("/", handler{ctx, root}.serve)
	s.Post("/fossils", handler{ctx, fossilize}.serve)

	// Launch result workers.
	rc := make(chan *Result)
	a.AddResultChan(rc)
	for i := 0; i < c.NumResultWorkers; i++ {
		go handleResults(rc)
	}

	return s
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

// HTTP handle to fossilize data.
func fossilize(w http.ResponseWriter, r *http.Request, p httprouter.Params, c *context) (interface{}, error) {
	data, url, err := parseFossilizeValues(r, c)
	if err != nil {
		return nil, err
	}

	if err := c.adapter.Fossilize(data, []byte(url)); err != nil {
		return nil, err
	}

	return "ok", nil
}

// Handles fossilization results
func handleResults(resultChan chan *Result) {
	for {
		r := <-resultChan

		body, err := json.Marshal(r.Evidence)
		if err != nil {
			log.Println(err)
			continue
		}

		url := string(r.Meta)
		res, err := http.Post(url, "application/json", bytes.NewReader(body))

		if err != nil {
			log.Println(err)
		} else if res.StatusCode >= 300 {
			log.Printf("%s: %d\n", url, res.StatusCode)
		}
	}
}

// Parses the data and callback URL from a request.
func parseFossilizeValues(r *http.Request, c *context) ([]byte, string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, "", err
	}

	datastr := r.Form.Get("data")
	if datastr == "" {
		return nil, "", &ErrData
	}

	l := len(datastr)
	if l < c.config.MinDataLen {
		return nil, "", &ErrDataLen
	}
	if c.config.MaxDataLen > 0 && l > c.config.MaxDataLen {
		return nil, "", &ErrDataLen
	}

	data, err := hex.DecodeString(datastr)
	if err != nil {
		return nil, "", &jsonhttp.ErrHTTP{err.Error(), 400}
	}

	url := r.Form.Get("callbackUrl")
	if url == "" {
		return nil, "", &ErrCallbackURL
	}

	return data, url, nil
}
