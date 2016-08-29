// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package storetestcases defines test cases to test stores.
package storetestcases

import (
	"testing"

	"github.com/stratumn/go/store"
)

// Factory wraps functions to allocate and free an adapter,
// and is used to run the tests on an adapter.
type Factory struct {
	// New creates an adapter.
	New func() (store.Adapter, error)

	// Free is an optional function to free an adapter.
	Free func(adapter store.Adapter)
}

// RunTests runs all the tests.
func (f Factory) RunTests(t *testing.T) {
	t.Run("DeleteSegment", f.TestDeleteSegment)
	t.Run("DeleteSegment_notFound", f.TestDeleteSegment_notFound)
	t.Run("FindSegments", f.TestFindSegments)
	t.Run("FindSegments_pagination", f.TestFindSegments_pagination)
	t.Run("FindSegment_empty", f.TestFindSegment_empty)
	t.Run("FindSegments_singleTag", f.TestFindSegments_singleTag)
	t.Run("FindSegments_multipleTags", f.TestFindSegments_multipleTags)
	t.Run("FindSegmentsMapID", f.TestFindSegmentsMapID)
	t.Run("FindSegmentsMapID_notFound", f.TestFindSegmentsMapID_notFound)
	t.Run("FindSegments_prevLinkHash", f.TestFindSegments_prevLinkHash)
	t.Run("FindSegments_prevLinkHashNotFound", f.TestFindSegments_prevLinkHashNotFound)
	t.Run("GetInfo", f.TestGetInfo)
	t.Run("GetMapIDs", f.TestGetMapIDs)
	t.Run("GetMapIDs_pagination", f.TestGetMapIDs_pagination)
	t.Run("GetMapIDs_empty", f.TestGetMapIDs_empty)
	t.Run("GetSegment", f.TestGetSegment)
	t.Run("GetSegment_updatedState", f.TestGetSegment_updatedState)
	t.Run("GetSegment_updatedMapID", f.TestGetSegment_updatedMapID)
	t.Run("GetSegment_notFound", f.TestGetSegment_notFound)
	t.Run("SaveSegment", f.TestSaveSegment)
	t.Run("SaveSegment_updatedState", f.TestSaveSegment_updatedState)
	t.Run("SaveSegment_updatedMapID", f.TestSaveSegment_updatedMapID)
	t.Run("SaveSegment_branch", f.TestSaveSegment_branch)
}

// RunBenchmarks runs all the benchmarks.
func (f Factory) RunBenchmarks(b *testing.B) {
	b.Run("DeleteSegment", f.BenchmarkDeleteSegment)
	b.Run("DeleteSegment_parallel", f.BenchmarkDeleteSegment_parallel)
	b.Run("GetSegment", f.BenchmarkGetSegment)
	b.Run("GetSegment_parallel", f.BenchmarkGetSegment_parallel)
	b.Run("SaveSegment", f.BenchmarkSaveSegment)
	b.Run("SaveSegment_parallel", f.BenchmarkSaveSegment_parallel)
	b.Run("SaveSegment_updatedState", f.BenchmarkSaveSegment_updatedState)
	b.Run("SaveSegment_updatedStateParallel", f.BenchmarkSaveSegment_updatedStateParallel)
	b.Run("SaveSegment_updatedMapID", f.BenchmarkSaveSegment_updatedMapID)
	b.Run("SaveSegment_updatedMapIDParallel", f.BenchmarkSaveSegment_updatedMapIDParallel)
}

func (f Factory) free(adapter store.Adapter) {
	if f.Free != nil {
		f.Free(adapter)
	}
}
