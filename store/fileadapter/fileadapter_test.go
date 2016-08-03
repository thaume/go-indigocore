package fileadapter

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stratumn/go/store/adapter/adaptertest"
)

func createAdapter(tb testing.TB) *FileAdapter {
	path, err := ioutil.TempDir("", "filestore")
	if err != nil {
		tb.Fatal(err)
	}

	return New(&Config{Path: path})
}

func TestGetInfo(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetInfo(t, a)
}

func TestSaveSegmentNew(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestSaveSegmentNew(t, a)
}

func TestSaveSegmentUpdateState(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestSaveSegmentUpdateState(t, a)
}

func TestSaveSegmentUpdateMapID(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestSaveSegmentUpdateMapID(t, a)
}

func TestSaveSegmentBranch(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestSaveSegmentBranch(t, a)
}

func TestGetSegmentFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetSegmentFound(t, a)
}

func TestGetSegmentUpdatedState(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetSegmentUpdatedState(t, a)
}

func TestGetSegmentUpdatedMapID(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetSegmentUpdatedMapID(t, a)
}

func TestGetSegmentNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetSegmentNotFound(t, a)
}

func TestDeleteSegmentFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestDeleteSegmentFound(t, a)
}

func TestDeleteSegmentNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestDeleteSegmentNotFound(t, a)
}

func TestFindSegmentsAll(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsAll(t, a)
}

func TestFindSegmentsPagination(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsPagination(t, a)
}

func TestFindSegmentsEmpty(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsEmpty(t, a)
}

func TestFindSegmentsSingleTag(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsSingleTag(t, a)
}

func TestFindSegmentsMultipleTags(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsMultipleTags(t, a)
}

func TestFindSegmentsMapIDFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsMapIDFound(t, a)
}

func TestFindSegmentsMapIDNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsMapIDNotFound(t, a)
}

func TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsPrevLinkHashFound(t, a)
}

func TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestFindSegmentsPrevLinkHashNotFound(t, a)
}

func TestGetMapIDsAll(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetMapIDsAll(t, a)
}

func TestGetMapIDsPagination(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetMapIDsPagination(t, a)
}

func TestGetMapIDsEmpty(t *testing.T) {
	a := createAdapter(t)
	defer os.RemoveAll(a.config.Path)
	adaptertest.TestGetMapIDsEmpty(t, a)
}

func BenchmarkSaveSegmentNew(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkSaveSegmentNew(b, a)
}

func BenchmarkSaveSegmentNewParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkSaveSegmentNewParallel(b, a)
}

func BenchmarkSaveSegmentUpdateState(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkSaveSegmentUpdateState(b, a)
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkSaveSegmentUpdateStateParallel(b, a)
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkSaveSegmentUpdateMapID(b, a)
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkSaveSegmentUpdateMapIDParallel(b, a)
}

func BenchmarkGetSegmentFound(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkGetSegmentFound(b, a)
}

func BenchmarkGetSegmentFoundParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkGetSegmentFoundParallel(b, a)
}

func BenchmarkDeleteSegmentFound(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkDeleteSegmentFound(b, a)
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
	a := createAdapter(b)
	defer os.RemoveAll(a.config.Path)
	adaptertest.BenchmarkDeleteSegmentFoundParallel(b, a)
}
