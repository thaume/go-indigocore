// Provides functionality to create an HTTP server from an adapter.
package httpserver

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	. "github.com/stratumn/go/fossilizer/adapter"
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
	Port             string
	CertFile         string
	KeyFile          string
	NumResultWorkers int
	MinDataLen       int
	MaxDataLen       int
	Verbose          bool
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
	if config.NumResultWorkers < 1 {
		config.NumResultWorkers = DEFAULT_NUM_RESULT_WORKERS
	}

	resultChan := make(chan *Result)
	adapter.AddResultChan(resultChan)

	// Launch result workers.
	for i := 0; i < config.NumResultWorkers; i++ {
		go HandleResults(resultChan)
	}

	r := httprouter.New()
	c := Context{adapter, config}

	r.NotFound = NotFoundHandler{JSONHandler{&c, NotFound}}

	r.GET("/", JSONHandler{&c, Root}.ServeHTTP)
	r.POST("/fossils", JSONHandler{&c, Fossilize}.ServeHTTP)

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

// HTTP handle to fossilize data.
func Fossilize(w http.ResponseWriter, r *http.Request, c *Context, p httprouter.Params) (interface{}, error) {
	data, callbackURL, err := parseFossilizeValues(r, c)
	if err != nil {
		return nil, err
	}

	if err := c.Adapter.Fossilize(data, []byte(callbackURL)); err != nil {
		return nil, err
	}

	return "ok", nil
}

// Handles fossilization results
func HandleResults(resultChan chan *Result) {
	for {
		result := <-resultChan

		body, err := json.Marshal(result.Evidence)
		if err != nil {
			log.Println(err)
			continue
		}

		callbackURL := string(result.Meta)
		res, err := http.Post(callbackURL, "application/json", bytes.NewReader(body))

		if err != nil {
			log.Println(err)
		} else if res.StatusCode >= 300 {
			log.Printf("%s: %d\n", callbackURL, res.StatusCode)
		}
	}
}

// Parses the data and callback URL from a request.
func parseFossilizeValues(r *http.Request, c *Context) ([]byte, string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, "", err
	}

	datastr := r.Form.Get("data")
	if datastr == "" {
		return nil, "", &ErrData
	}

	length := len(datastr)
	if length < c.Config.MinDataLen {
		return nil, "", &ErrDataLen
	}
	if c.Config.MaxDataLen > 0 && length > c.Config.MaxDataLen {
		return nil, "", &ErrDataLen
	}

	data, err := hex.DecodeString(datastr)
	if err != nil {
		return nil, "", &ErrHTTP{err.Error(), 400}
	}

	callbackURL := r.Form.Get("callbackUrl")
	if callbackURL == "" {
		return nil, "", &ErrCallbackURL
	}

	return data, callbackURL, nil
}
