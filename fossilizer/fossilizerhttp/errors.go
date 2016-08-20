// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizerhttp

import (
	"github.com/stratumn/go/jsonhttp"
)

func newErrData(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "data required"
	}
	return jsonhttp.NewErrBadRequest(msg)
}

func newErrDataLen(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "invalid data length"
	}
	return jsonhttp.NewErrBadRequest(msg)
}

func newErrCallbackURL(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "callback URL required"
	}
	return jsonhttp.NewErrBadRequest(msg)
}
