// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package storetestcases defines test cases to test stores.
package storetestcases

import (
	"testing"

	"github.com/stratumn/go/store"
)

// Factory contains function to allocate and free an adapter.
type Factory struct {
	// New create an adapter.
	New func() (store.Adapter, error)

	// Free is an optional function to free an adapter.
	Free func(adapter store.Adapter)
}

// RunTests runs all the tests.
func (f Factory) RunTests(t *testing.T) {
	t.Run("DeleteSegmentFound", f.TestDeleteSegmentFound)
	t.Run("DeleteSegmentNotFound", f.TestDeleteSegmentNotFound)
	t.Run("FindSegmentsAll", f.TestFindSegmentsAll)
	t.Run("FindSegmentsPagination", f.TestFindSegmentsPagination)
	t.Run("FindSegmentsEmpty", f.TestFindSegmentsEmpty)
	t.Run("FindSegmentsSingleTag", f.TestFindSegmentsSingleTag)
	t.Run("FindSegmentsMultipleTags", f.TestFindSegmentsMultipleTags)
	t.Run("FindSegmentsMapIDFound", f.TestFindSegmentsMapIDFound)
	t.Run("FindSegmentsMapIDNotFound", f.TestFindSegmentsMapIDNotFound)
	t.Run("FindSegmentsPrevLinkHashFound", f.TestFindSegmentsPrevLinkHashFound)
	t.Run("FindSegmentsPrevLinkHashNotFound", f.TestFindSegmentsPrevLinkHashNotFound)
	t.Run("GetInfo", f.TestGetInfo)
	t.Run("GetMapIDsAll", f.TestGetMapIDsAll)
	t.Run("GetMapIDsPagination", f.TestGetMapIDsPagination)
	t.Run("GetMapIDsEmpty", f.TestGetMapIDsEmpty)
	t.Run("GetSegmentFound", f.TestGetSegmentFound)
	t.Run("GetSegmentUpdatedState", f.TestGetSegmentUpdatedState)
	t.Run("GetSegmentUpdatedMapID", f.TestGetSegmentUpdatedMapID)
	t.Run("GetSegmentNotFound", f.TestGetSegmentNotFound)
	t.Run("SaveSegmentNew", f.TestSaveSegmentNew)
	t.Run("SaveSegmentUpdateState", f.TestSaveSegmentUpdateState)
	t.Run("SaveSegmentUpdateMapID", f.TestSaveSegmentUpdateMapID)
	t.Run("SaveSegmentBranch", f.TestSaveSegmentBranch)
}

// RunBenchmarks runs all the benchmarks.
func (f Factory) RunBenchmarks(b *testing.B) {
	b.Run("DeleteSegmentFound", f.BenchmarkDeleteSegmentFound)
	b.Run("DeleteSegmentFoundParallel", f.BenchmarkDeleteSegmentFoundParallel)
	b.Run("GetSegmentFound", f.BenchmarkGetSegmentFound)
	b.Run("GetSegmentFoundParallel", f.BenchmarkGetSegmentFoundParallel)
	b.Run("SaveSegmentNew", f.BenchmarkSaveSegmentNew)
	b.Run("SaveSegmentNewParallel", f.BenchmarkSaveSegmentNewParallel)
	b.Run("SaveSegmentUpdateState", f.BenchmarkSaveSegmentUpdateState)
	b.Run("SaveSegmentUpdateStateParallel", f.BenchmarkSaveSegmentUpdateStateParallel)
	b.Run("SaveSegmentUpdateMapID", f.BenchmarkSaveSegmentUpdateMapID)
	b.Run("SaveSegmentUpdateMapIDParallel", f.BenchmarkSaveSegmentUpdateMapIDParallel)
}

func (f Factory) free(adapter store.Adapter) {
	if f.Free != nil {
		f.Free(adapter)
	}
}
