// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizerhttp

import (
	"net/http"

	"github.com/stratumn/go/jsonhttp"
)

var (
	// ErrData is an error for when no data is given to fossilize.
	ErrData = jsonhttp.NewErrHTTP("data required", http.StatusBadRequest)

	// ErrDataLen is an error for the data given to fossilize is either too short or too long.
	ErrDataLen = jsonhttp.NewErrHTTP("invalid data length", http.StatusBadRequest)

	// ErrCallbackURL is an error for when no callback URL is given to fossilize.
	ErrCallbackURL = jsonhttp.NewErrHTTP("callback URL required", http.StatusBadRequest)
)
