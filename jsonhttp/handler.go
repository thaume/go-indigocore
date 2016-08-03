package jsonhttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// A JSON handle function.
type JSONHandle func(http.ResponseWriter, *http.Request, httprouter.Params, *Config) (interface{}, error)

// A JSON handler.
type JSONHandler struct {
	config *Config
	serve  JSONHandle
}

// Implements http.Server.
func (h JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	data, err := h.serve(w, r, p, h.config)

	if err != nil {
		JSONError(w, err, h.config)
		return
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "unexpected error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// A not found handler.
type NotFoundHandler struct {
	JSONHandler
}

// Implements HTTP server.
func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.JSONHandler.ServeHTTP(w, r, nil)
}

// Renders an error as JSON.
func JSONError(w http.ResponseWriter, err error, c *Config) {
	e, ok := err.(*ErrHTTP)

	if !ok {
		log.Println(err.Error())
		e = &ErrInternalServer
	} else if c.Verbose {
		log.Println(err.Error())
	}

	js := e.JSONEncode()

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(js), e.Status)
}
