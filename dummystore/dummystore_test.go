// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

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

func BenchmarkDummystore(b *testing.B) {
	storetestcases.Factory{
		New: func() (store.Adapter, error) {
			return New(&Config{}), nil
		},
	}.RunBenchmarks(b)
}
