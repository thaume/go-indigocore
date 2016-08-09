// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storehttp

import (
	"github.com/stratumn/go/jsonhttp"
)

func newErrOffset(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "offset must be a positive integer"
	}

	return jsonhttp.NewErrBadRequest(msg)
}

func newErrLimit(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "limit must be a posive integer"
	}

	return jsonhttp.NewErrBadRequest(msg)
}
