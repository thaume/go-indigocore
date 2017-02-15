// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package storehttp

import (
	"fmt"

	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store"
)

func newErrOffset(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "offset must be a positive integer"
	}
	return jsonhttp.NewErrBadRequest(msg)
}

func newErrLimit(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = fmt.Sprintf("limit must be a posive integer less than or equal to %d", store.MaxLimit)
	}
	return jsonhttp.NewErrBadRequest(msg)
}

func newErrPrevLinkHash(msg string) jsonhttp.ErrHTTP {
	if msg == "" {
		msg = "prevLinkHash must be a 64 byte long hexadecimal string"
	}
	return jsonhttp.NewErrBadRequest(msg)
}
