// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package filestore

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stratumn/go/store/storetesting"
)

func createAdapter(tb testing.TB) *FileStore {
	path, err := ioutil.TempDir("", "filestore")
	if err != nil {
		tb.Fatal(err)
	}

	return New(&Config{Path: path})
}

func TestGetInfo(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetInfo(t, a)
}

func TestSaveSegmentNew(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestSaveSegmentNew(t, a)
}

func TestSaveSegmentUpdateState(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestSaveSegmentUpdateState(t, a)
}

func TestSaveSegmentUpdateMapID(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestSaveSegmentUpdateMapID(t, a)
}

func TestSaveSegmentBranch(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestSaveSegmentBranch(t, a)
}

func TestGetSegmentFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetSegmentFound(t, a)
}

func TestGetSegmentUpdatedState(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetSegmentUpdatedState(t, a)
}

func TestGetSegmentUpdatedMapID(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetSegmentUpdatedMapID(t, a)
}

func TestGetSegmentNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetSegmentNotFound(t, a)
}

func TestDeleteSegmentFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestDeleteSegmentFound(t, a)
}

func TestDeleteSegmentNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestDeleteSegmentNotFound(t, a)
}

func TestFindSegmentsAll(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsAll(t, a)
}

func TestFindSegmentsPagination(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsPagination(t, a)
}

func TestFindSegmentsEmpty(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsEmpty(t, a)
}

func TestFindSegmentsSingleTag(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsSingleTag(t, a)
}

func TestFindSegmentsMultipleTags(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsMultipleTags(t, a)
}

func TestFindSegmentsMapIDFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsMapIDFound(t, a)
}

func TestFindSegmentsMapIDNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsMapIDNotFound(t, a)
}

func TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsPrevLinkHashFound(t, a)
}

func TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestFindSegmentsPrevLinkHashNotFound(t, a)
}

func TestGetMapIDsAll(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetMapIDsAll(t, a)
}

func TestGetMapIDsPagination(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetMapIDsPagination(t, a)
}

func TestGetMapIDsEmpty(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetesting.TestGetMapIDsEmpty(t, a)
}

func BenchmarkSaveSegmentNew(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkSaveSegmentNew(b, a)
}

func BenchmarkSaveSegmentNewParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkSaveSegmentNewParallel(b, a)
}

func BenchmarkSaveSegmentUpdateState(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkSaveSegmentUpdateState(b, a)
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkSaveSegmentUpdateStateParallel(b, a)
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkSaveSegmentUpdateMapID(b, a)
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkSaveSegmentUpdateMapIDParallel(b, a)
}

func BenchmarkGetSegmentFound(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkGetSegmentFound(b, a)
}

func BenchmarkGetSegmentFoundParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkGetSegmentFoundParallel(b, a)
}

func BenchmarkDeleteSegmentFound(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkDeleteSegmentFound(b, a)
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetesting.BenchmarkDeleteSegmentFoundParallel(b, a)
}
