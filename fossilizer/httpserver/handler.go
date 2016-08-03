package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/stratumn/go/fossilizer/adapter"
)

// A handler context.
type Context struct {
	Adapter adapter.Adapter
	Config  *Config
}

// A JSON handle function.
type JSONHandle func(http.ResponseWriter, *http.Request, *Context, httprouter.Params) (interface{}, error)

// A JSON handler.
type JSONHandler struct {
	Context *Context
	Serve   JSONHandle
}

// Implements HTTP server.
func (h JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error

	data, err := h.Serve(w, r, h.Context, p)

	if err != nil {
		JSONError(h.Context, w, err)
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

// Implements HTTP handler.
func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.JSONHandler.ServeHTTP(w, r, nil)
}

// Renders an error as JSON.
func JSONError(c *Context, w http.ResponseWriter, err error) {
	errHTTP, ok := err.(*ErrHTTP)

	if !ok {
		log.Println(err.Error())
		errHTTP = &ErrInternalServer
	} else if c.Config.Verbose {
		log.Println(err.Error())
	}

	js := errHTTP.JSONEncode()

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, string(js), errHTTP.Status)
}
