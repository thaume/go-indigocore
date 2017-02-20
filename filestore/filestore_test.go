// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package filestore

import (
	"os"
	"testing"

	"github.com/stratumn/sdk/store"
	"github.com/stratumn/sdk/store/storetestcases"
)

func TestFilestore(t *testing.T) {
	storetestcases.Factory{
		New: func() (store.Adapter, error) {
			return createAdapter(t), nil
		},
		Free: func(s store.Adapter) {
			a := s.(*FileStore)
			defer os.RemoveAll(a.config.Path)
		},
	}.RunTests(t)
}
