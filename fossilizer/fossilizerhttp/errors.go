// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
