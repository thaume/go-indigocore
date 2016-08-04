// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizerhttp

import (
	"github.com/stratumn/go/jsonhttp"
)

var (
	// ErrData is an error for when no data is given to fossilize.
	ErrData = jsonhttp.ErrHTTP{Msg: "data required", Status: 400}

	// ErrDataLen is an error for the data given to fossilize is either too short or too long.
	ErrDataLen = jsonhttp.ErrHTTP{Msg: "invalid data length", Status: 400}

	// ErrCallbackURL is an error for when no callback URL is given to fossilize.
	ErrCallbackURL = jsonhttp.ErrHTTP{Msg: "callback URL required", Status: 400}
)
