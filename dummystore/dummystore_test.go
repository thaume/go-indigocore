// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package dummystore

import (
	"testing"

	"github.com/stratumn/go/store/storetestcases"
)

func TestGetInfo(t *testing.T) {
	storetestcases.TestGetInfo(t, New(""))
}

func TestSaveSegmentNew(t *testing.T) {
	storetestcases.TestSaveSegmentNew(t, New(""))
}

func TestSaveSegmentUpdateState(t *testing.T) {
	storetestcases.TestSaveSegmentUpdateState(t, New(""))
}

func TestSaveSegmentUpdateMapID(t *testing.T) {
	storetestcases.TestSaveSegmentUpdateMapID(t, New(""))
}

func TestSaveSegmentBranch(t *testing.T) {
	storetestcases.TestSaveSegmentBranch(t, New(""))
}

func TestGetSegmentFound(t *testing.T) {
	storetestcases.TestGetSegmentFound(t, New(""))
}

func TestGetSegmentUpdatedState(t *testing.T) {
	storetestcases.TestGetSegmentUpdatedState(t, New(""))
}

func TestGetSegmentUpdatedMapID(t *testing.T) {
	storetestcases.TestGetSegmentUpdatedMapID(t, New(""))
}

func TestGetSegmentNotFound(t *testing.T) {
	storetestcases.TestGetSegmentNotFound(t, New(""))
}

func TestDeleteSegmentFound(t *testing.T) {
	storetestcases.TestDeleteSegmentFound(t, New(""))
}

func TestDeleteSegmentNotFound(t *testing.T) {
	storetestcases.TestDeleteSegmentNotFound(t, New(""))
}

func TestFindSegmentsAll(t *testing.T) {
	storetestcases.TestFindSegmentsAll(t, New(""))
}

func TestFindSegmentsPagination(t *testing.T) {
	storetestcases.TestFindSegmentsPagination(t, New(""))
}

func TestFindSegmentsEmpty(t *testing.T) {
	storetestcases.TestFindSegmentsEmpty(t, New(""))
}

func TestFindSegmentsSingleTag(t *testing.T) {
	storetestcases.TestFindSegmentsSingleTag(t, New(""))
}

func TestFindSegmentsMultipleTags(t *testing.T) {
	storetestcases.TestFindSegmentsMultipleTags(t, New(""))
}

func TestFindSegmentsMapIDFound(t *testing.T) {
	storetestcases.TestFindSegmentsMapIDFound(t, New(""))
}

func TestFindSegmentsMapIDNotFound(t *testing.T) {
	storetestcases.TestFindSegmentsMapIDNotFound(t, New(""))
}

func TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	storetestcases.TestFindSegmentsPrevLinkHashFound(t, New(""))
}

func TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	storetestcases.TestFindSegmentsPrevLinkHashNotFound(t, New(""))
}

func TestGetMapIDsAll(t *testing.T) {
	storetestcases.TestGetMapIDsAll(t, New(""))
}

func TestGetMapIDsPagination(t *testing.T) {
	storetestcases.TestGetMapIDsPagination(t, New(""))
}

func TestGetMapIDsEmpty(t *testing.T) {
	storetestcases.TestGetMapIDsEmpty(t, New(""))
}

func BenchmarkSaveSegmentNew(b *testing.B) {
	storetestcases.BenchmarkSaveSegmentNew(b, New(""))
}

func BenchmarkSaveSegmentNewParallel(b *testing.B) {
	storetestcases.BenchmarkSaveSegmentNewParallel(b, New(""))
}

func BenchmarkSaveSegmentUpdateState(b *testing.B) {
	storetestcases.BenchmarkSaveSegmentUpdateState(b, New(""))
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	storetestcases.BenchmarkSaveSegmentUpdateStateParallel(b, New(""))
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	storetestcases.BenchmarkSaveSegmentUpdateMapID(b, New(""))
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	storetestcases.BenchmarkSaveSegmentUpdateMapIDParallel(b, New(""))
}

func BenchmarkGetSegmentFound(b *testing.B) {
	storetestcases.BenchmarkGetSegmentFound(b, New(""))
}

func BenchmarkGetSegmentFoundParallel(b *testing.B) {
	storetestcases.BenchmarkGetSegmentFoundParallel(b, New(""))
}

func BenchmarkDeleteSegmentFound(b *testing.B) {
	storetestcases.BenchmarkDeleteSegmentFound(b, New(""))
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
	storetestcases.BenchmarkDeleteSegmentFoundParallel(b, New(""))
}
