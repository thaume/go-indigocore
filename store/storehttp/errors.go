// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storehttp

import (
	"github.com/stratumn/go/jsonhttp"
)

// NewErrOffset creates an error for when an invalid pagination offset is given.
// If the message is empty, the default is "offset must be a positive integer".
func NewErrOffset(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "offset must be a positive integer"
	}

	return jsonhttp.NewErrBadRequest(msg)
}

// NewErrLimit creates an error for when an invalid pagination offset is given.
// If the message is empty, the default is "limit must be a posive integer".
func NewErrLimit(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "limit must be a posive integer"
	}

	return jsonhttp.NewErrBadRequest(msg)
}
