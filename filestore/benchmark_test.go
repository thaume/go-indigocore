// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package filestore

import (
	"os"
	"testing"

	"github.com/stratumn/go/store"
	"github.com/stratumn/go/store/storetestcases"
)

func BenchmarkFilestore(b *testing.B) {
	storetestcases.Factory{
		New: func() (store.Adapter, error) {
			return createAdapter(b), nil
		},
		Free: func(s store.Adapter) {
			a := s.(*FileStore)
			defer os.RemoveAll(a.config.Path)
		},
	}.RunBenchmarks(b)
}
