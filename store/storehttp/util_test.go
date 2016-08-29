// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storehttp

import (
	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/storetesting"
)

func createServer() (*jsonhttp.Server, *storetesting.MockAdapter) {
	a := &storetesting.MockAdapter{}
	s := New(a, &jsonhttp.Config{})

	return s, a
}
