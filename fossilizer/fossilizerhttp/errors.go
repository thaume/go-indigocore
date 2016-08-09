// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizerhttp

import (
	"github.com/stratumn/go/jsonhttp"
)

// NewErrData creates an error for when no data is given to fossilize.
// If the message is empty, the default is "data required".
func NewErrData(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "data required"
	}

	return jsonhttp.NewErrBadRequest(msg)
}

// NewErrDataLen creates an error for when the data given to fossilize is either too short or too long.
// If the message is empty, the default is "invalid data length".
func NewErrDataLen(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "invalid data length"
	}

	return jsonhttp.NewErrBadRequest(msg)
}

// NewErrCallbackURL creates an error for when no callback URL is given.
// If the message is empty, the default is "callback URL required".
func NewErrCallbackURL(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "callback URL required"
	}

	return jsonhttp.NewErrBadRequest(msg)
}
