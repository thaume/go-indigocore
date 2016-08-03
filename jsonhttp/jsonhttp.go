// A package to create simple HTTP servers that render JSON.
package jsonhttp

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	router *httprouter.Router
	config *Config
}

// Implements HTTP server.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Makes a new server.
func New(config *Config) *Server {
	r := httprouter.New()
	r.NotFound = NotFoundHandler{JSONHandler{config, NotFound}}
	return &Server{r, config}
}

// Adds a GET route.
func (s *Server) Get(path string, handle JSONHandle) {
	s.router.GET(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Adds a POST route.
func (s *Server) Post(path string, handle JSONHandle) {
	s.router.POST(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Adds a PUT route.
func (s *Server) Put(path string, handle JSONHandle) {
	s.router.PUT(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Adds a DELETE route.
func (s *Server) Delete(path string, handle JSONHandle) {
	s.router.DELETE(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Adds a PATCH route.
func (s *Server) Patch(path string, handle JSONHandle) {
	s.router.PATCH(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Adds a HEAD route.
func (s *Server) Head(path string, handle JSONHandle) {
	s.router.HEAD(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Adds an OPTIONS route.
func (s *Server) Options(path string, handle JSONHandle) {
	s.router.OPTIONS(path, JSONHandler{s.config, handle}.ServeHTTP)
}

// Starts listening and serving.
func (s *Server) ListenAndServe() error {
	p := s.config.Port

	if p == "" {
		p = DEFAULT_PORT
	}

	if s.config.CertFile != "" && s.config.KeyFile != "" {
		return http.ListenAndServeTLS(p, s.config.CertFile, s.config.KeyFile, s)
	}

	return http.ListenAndServe(p, s)
}

// HTTP handle for 404s.
func NotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params, _ *Config) (interface{}, error) {
	return nil, &ErrNotFound
}
