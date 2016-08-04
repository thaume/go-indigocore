// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package dummystore

import (
	"testing"

	"github.com/stratumn/go/store/storetesting"
)

func TestGetInfo(t *testing.T) {
	storetesting.TestGetInfo(t, New(""))
}

func TestSaveSegmentNew(t *testing.T) {
	storetesting.TestSaveSegmentNew(t, New(""))
}

func TestSaveSegmentUpdateState(t *testing.T) {
	storetesting.TestSaveSegmentUpdateState(t, New(""))
}

func TestSaveSegmentUpdateMapID(t *testing.T) {
	storetesting.TestSaveSegmentUpdateMapID(t, New(""))
}

func TestSaveSegmentBranch(t *testing.T) {
	storetesting.TestSaveSegmentBranch(t, New(""))
}

func TestGetSegmentFound(t *testing.T) {
	storetesting.TestGetSegmentFound(t, New(""))
}

func TestGetSegmentUpdatedState(t *testing.T) {
	storetesting.TestGetSegmentUpdatedState(t, New(""))
}

func TestGetSegmentUpdatedMapID(t *testing.T) {
	storetesting.TestGetSegmentUpdatedMapID(t, New(""))
}

func TestGetSegmentNotFound(t *testing.T) {
	storetesting.TestGetSegmentNotFound(t, New(""))
}

func TestDeleteSegmentFound(t *testing.T) {
	storetesting.TestDeleteSegmentFound(t, New(""))
}

func TestDeleteSegmentNotFound(t *testing.T) {
	storetesting.TestDeleteSegmentNotFound(t, New(""))
}

func TestFindSegmentsAll(t *testing.T) {
	storetesting.TestFindSegmentsAll(t, New(""))
}

func TestFindSegmentsPagination(t *testing.T) {
	storetesting.TestFindSegmentsPagination(t, New(""))
}

func TestFindSegmentsEmpty(t *testing.T) {
	storetesting.TestFindSegmentsEmpty(t, New(""))
}

func TestFindSegmentsSingleTag(t *testing.T) {
	storetesting.TestFindSegmentsSingleTag(t, New(""))
}

func TestFindSegmentsMultipleTags(t *testing.T) {
	storetesting.TestFindSegmentsMultipleTags(t, New(""))
}

func TestFindSegmentsMapIDFound(t *testing.T) {
	storetesting.TestFindSegmentsMapIDFound(t, New(""))
}

func TestFindSegmentsMapIDNotFound(t *testing.T) {
	storetesting.TestFindSegmentsMapIDNotFound(t, New(""))
}

func TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	storetesting.TestFindSegmentsPrevLinkHashFound(t, New(""))
}

func TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	storetesting.TestFindSegmentsPrevLinkHashNotFound(t, New(""))
}

func TestGetMapIDsAll(t *testing.T) {
	storetesting.TestGetMapIDsAll(t, New(""))
}

func TestGetMapIDsPagination(t *testing.T) {
	storetesting.TestGetMapIDsPagination(t, New(""))
}

func TestGetMapIDsEmpty(t *testing.T) {
	storetesting.TestGetMapIDsEmpty(t, New(""))
}

func BenchmarkSaveSegmentNew(b *testing.B) {
	storetesting.BenchmarkSaveSegmentNew(b, New(""))
}

func BenchmarkSaveSegmentNewParallel(b *testing.B) {
	storetesting.BenchmarkSaveSegmentNewParallel(b, New(""))
}

func BenchmarkSaveSegmentUpdateState(b *testing.B) {
	storetesting.BenchmarkSaveSegmentUpdateState(b, New(""))
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	storetesting.BenchmarkSaveSegmentUpdateStateParallel(b, New(""))
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	storetesting.BenchmarkSaveSegmentUpdateMapID(b, New(""))
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	storetesting.BenchmarkSaveSegmentUpdateMapIDParallel(b, New(""))
}

func BenchmarkGetSegmentFound(b *testing.B) {
	storetesting.BenchmarkGetSegmentFound(b, New(""))
}

func BenchmarkGetSegmentFoundParallel(b *testing.B) {
	storetesting.BenchmarkGetSegmentFoundParallel(b, New(""))
}

func BenchmarkDeleteSegmentFound(b *testing.B) {
	storetesting.BenchmarkDeleteSegmentFound(b, New(""))
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
	storetesting.BenchmarkDeleteSegmentFoundParallel(b, New(""))
}
