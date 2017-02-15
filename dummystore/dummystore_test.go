// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package dummystore

import (
	"testing"

	"github.com/stratumn/go/store"
	"github.com/stratumn/go/store/storetestcases"
)

func TestDummystore(t *testing.T) {
	storetestcases.Factory{
		New: func() (store.Adapter, error) {
			return New(&Config{}), nil
		},
	}.RunTests(t)
}
