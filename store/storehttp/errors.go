package storehttp

import (
	"github.com/stratumn/go/jsonhttp"
)

var (
	// ErrOffset is an error for when an invalid pagination offset is given.
	ErrOffset = jsonhttp.ErrHTTP{Msg: "offset must be a positive integer", Status: 400}

	// ErrLimit is an error for when an invalid pagination limit is given.
	ErrLimit = jsonhttp.ErrHTTP{Msg: "limit must be a posive integer", Status: 400}
)
