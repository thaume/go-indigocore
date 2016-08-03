// Package fossilizerhttp is used to create an HTTP server from a fossilizer adapter.
// It servers the following routes:
//	GET /
//		Renders information about the fossilizer.
//	POST /fossils
//		Requests data to be fossilized.
//		Form.data should be a hex encoded buffer.
//		Form.callbackUrl should be URL.
package fossilizerhttp

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/stratumn/go/fossilizer"
	"github.com/stratumn/go/jsonhttp"
)

const (
	// DefaultPort is the default port of the server.
	DefaultPort = ":6000"
	// DefaultNumResultWorkers is the default number of goroutines that will be used
	// to handle fossilizer results.
	DefaultNumResultWorkers = 8
	// DefaultMinDataLen is the default minimum fossilize data length.
	DefaultMinDataLen = 32
	// DefaultMaxDataLen is the default maximum fossilize data length.
	DefaultMaxDataLen = 64
	// DefaultVerbose is whether verbose output should be enabled by default.
	DefaultVerbose = false
)

// Config contains configuration options for the server.
type Config struct {
	jsonhttp.Config
	// The default number of goroutines that will be used to handle fossilizer results.
	NumResultWorkers int
	// The minimum fossilize data length.
	MinDataLen int
	// The maximum fossilize data length.
	MaxDataLen int
}

type context struct {
	adapter fossilizer.Adapter
	config  *Config
}

type handle func(http.ResponseWriter, *http.Request, httprouter.Params, *context) (interface{}, error)

type handler struct {
	context *context
	handle  handle
}

func (h handler) serve(w http.ResponseWriter, r *http.Request, p httprouter.Params, _ *jsonhttp.Config) (interface{}, error) {
	return h.handle(w, r, p, h.context)
}

// New create a new instance of a server.
func New(a fossilizer.Adapter, c *Config) *jsonhttp.Server {
	if c.NumResultWorkers < 1 {
		c.NumResultWorkers = DefaultNumResultWorkers
	}

	s := jsonhttp.New(&c.Config)
	ctx := &context{a, c}

	s.Get("/", handler{ctx, root}.serve)
	s.Post("/fossils", handler{ctx, fossilize}.serve)

	// Launch result workers.
	rc := make(chan *fossilizer.Result)
	a.AddResultChan(rc)
	for i := 0; i < c.NumResultWorkers; i++ {
		go handleResults(rc)
	}

	return s
}

func root(w http.ResponseWriter, r *http.Request, _ httprouter.Params, c *context) (interface{}, error) {
	info, err := c.adapter.GetInfo()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"adapter": info,
	}, nil
}

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

func handleResults(resultChan chan *fossilizer.Result) {
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
		return nil, "", &jsonhttp.ErrHTTP{Msg: err.Error(), Status: 400}
	}

	url := r.Form.Get("callbackUrl")
	if url == "" {
		return nil, "", &ErrCallbackURL
	}

	return data, url, nil
}
