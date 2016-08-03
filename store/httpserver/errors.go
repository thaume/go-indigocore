package httpserver

import (
	"github.com/stratumn/go/jsonhttp"
)

var (
	ErrOffset = jsonhttp.ErrHTTP{"offset must be a positive integer", 400}
	ErrLimit  = jsonhttp.ErrHTTP{"limit must be a posive integer", 400}
)
