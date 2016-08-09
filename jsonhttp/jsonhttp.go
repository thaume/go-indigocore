// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package jsonhttp defines a simple HTTP server that renders JSON.
//
// Routes can be added by passing a handle that should return JSON serializable data or an error.
package jsonhttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	// DefaultPort is the default port of the server.
	DefaultPort = ":5000"

	// DefaultVerbose is whether verbose output should be enabled by default.
	DefaultVerbose = false
)

// Config contains configuration options for the server.
type Config struct {
	// The port of the server.
	Port string

	// Optionally, the path to a TLS certificate.
	CertFile string

	// Optionally, the path to a TLS private key.
	KeyFile string

	// Whether to enable verbose output.
	Verbose bool
}

// Server is the type that implements net/http.Handler.
type Server struct {
	router *httprouter.Router
	config *Config
}

// JSONHandle is the function type for a route handle.
type JSONHandle func(http.ResponseWriter, *http.Request, httprouter.Params, *Config) (interface{}, error)

// NotFoundHandler is a handlers for undefined routes.
type NotFoundHandler struct {
	handler
}

// ServeHTTP implements net/http.Handler.ServeHTTP.
func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r, nil)
}

// ServeHTTP implements net/http.Handler.ServeHTTP.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// New creates an instance of Server.
func New(config *Config) *Server {
	r := httprouter.New()
	r.NotFound = NotFoundHandler{handler{config, notFound}}
	return &Server{r, config}
}

// Get adds a GET route.
func (s *Server) Get(path string, handle JSONHandle) {
	s.router.GET(path, handler{s.config, handle}.ServeHTTP)
}

// Post adds a POST route.
func (s *Server) Post(path string, handle JSONHandle) {
	s.router.POST(path, handler{s.config, handle}.ServeHTTP)
}

// Put adds PUT route.
func (s *Server) Put(path string, handle JSONHandle) {
	s.router.PUT(path, handler{s.config, handle}.ServeHTTP)
}

// Delete adds DELETE route.
func (s *Server) Delete(path string, handle JSONHandle) {
	s.router.DELETE(path, handler{s.config, handle}.ServeHTTP)
}

// Patch adds a PATCH route.
func (s *Server) Patch(path string, handle JSONHandle) {
	s.router.PATCH(path, handler{s.config, handle}.ServeHTTP)
}

// Head adds a HEAD route.
func (s *Server) Head(path string, handle JSONHandle) {
	s.router.HEAD(path, handler{s.config, handle}.ServeHTTP)
}

// Options adds an OPTIONS route.
func (s *Server) Options(path string, handle JSONHandle) {
	s.router.OPTIONS(path, handler{s.config, handle}.ServeHTTP)
}

// ListenAndServe starts the server.
func (s *Server) ListenAndServe() error {
	p := s.config.Port

	if p == "" {
		p = DefaultPort
	}

	if s.config.CertFile != "" && s.config.KeyFile != "" {
		return http.ListenAndServeTLS(p, s.config.CertFile, s.config.KeyFile, s)
	}

	return http.ListenAndServe(p, s)
}

func notFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params, _ *Config) (interface{}, error) {
	return nil, NewErrNotFound("")
}

type handler struct {
	config *Config
	serve  JSONHandle
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	data, err := h.serve(w, r, p, h.config)

	if err != nil {
		renderErr(w, err, h.config)
		return
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func renderErr(w http.ResponseWriter, err error, c *Config) {
	e, ok := err.(ErrHTTP)

	if !ok {
		log.Println(err.Error())
		e = NewErrInternalServer("")
	} else if c.Verbose {
		log.Println(err.Error())
	}

	js := e.JSONMarshal()

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(js), e.Status())
}
