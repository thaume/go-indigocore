package httpserver

import (
	"github.com/stratumn/go/jsonhttp"
)

var (
	ErrData        = jsonhttp.ErrHTTP{"data required", 400}
	ErrDataLen     = jsonhttp.ErrHTTP{"invalid data length", 400}
	ErrCallbackURL = jsonhttp.ErrHTTP{"callback URL required", 400}
)
