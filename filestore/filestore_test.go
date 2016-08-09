// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package filestore

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stratumn/go/store/storetestcases"
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
	storetestcases.TestGetInfo(t, a)
}

func TestSaveSegmentNew(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestSaveSegmentNew(t, a)
}

func TestSaveSegmentUpdateState(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestSaveSegmentUpdateState(t, a)
}

func TestSaveSegmentUpdateMapID(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestSaveSegmentUpdateMapID(t, a)
}

func TestSaveSegmentBranch(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestSaveSegmentBranch(t, a)
}

func TestGetSegmentFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetSegmentFound(t, a)
}

func TestGetSegmentUpdatedState(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetSegmentUpdatedState(t, a)
}

func TestGetSegmentUpdatedMapID(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetSegmentUpdatedMapID(t, a)
}

func TestGetSegmentNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetSegmentNotFound(t, a)
}

func TestDeleteSegmentFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestDeleteSegmentFound(t, a)
}

func TestDeleteSegmentNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestDeleteSegmentNotFound(t, a)
}

func TestFindSegmentsAll(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsAll(t, a)
}

func TestFindSegmentsPagination(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsPagination(t, a)
}

func TestFindSegmentsEmpty(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsEmpty(t, a)
}

func TestFindSegmentsSingleTag(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsSingleTag(t, a)
}

func TestFindSegmentsMultipleTags(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsMultipleTags(t, a)
}

func TestFindSegmentsMapIDFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsMapIDFound(t, a)
}

func TestFindSegmentsMapIDNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsMapIDNotFound(t, a)
}

func TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsPrevLinkHashFound(t, a)
}

func TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestFindSegmentsPrevLinkHashNotFound(t, a)
}

func TestGetMapIDsAll(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetMapIDsAll(t, a)
}

func TestGetMapIDsPagination(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetMapIDsPagination(t, a)
}

func TestGetMapIDsEmpty(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	storetestcases.TestGetMapIDsEmpty(t, a)
}

func BenchmarkSaveSegmentNew(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkSaveSegmentNew(b, a)
}

func BenchmarkSaveSegmentNewParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkSaveSegmentNewParallel(b, a)
}

func BenchmarkSaveSegmentUpdateState(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkSaveSegmentUpdateState(b, a)
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkSaveSegmentUpdateStateParallel(b, a)
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkSaveSegmentUpdateMapID(b, a)
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkSaveSegmentUpdateMapIDParallel(b, a)
}

func BenchmarkGetSegmentFound(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkGetSegmentFound(b, a)
}

func BenchmarkGetSegmentFoundParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkGetSegmentFoundParallel(b, a)
}

func BenchmarkDeleteSegmentFound(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkDeleteSegmentFound(b, a)
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	storetestcases.BenchmarkDeleteSegmentFoundParallel(b, a)
}
