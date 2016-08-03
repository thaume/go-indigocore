package dummyadapter

import (
	"testing"

	"github.com/stratumn/go/store/adapter/adaptertest"
)

func TestGetInfo(t *testing.T) {
	adaptertest.TestGetInfo(t, New(""))
}

func TestSaveSegmentNew(t *testing.T) {
	adaptertest.TestSaveSegmentNew(t, New(""))
}

func TestSaveSegmentUpdateState(t *testing.T) {
	adaptertest.TestSaveSegmentUpdateState(t, New(""))
}

func TestSaveSegmentUpdateMapID(t *testing.T) {
	adaptertest.TestSaveSegmentUpdateMapID(t, New(""))
}

func TestSaveSegmentBranch(t *testing.T) {
	adaptertest.TestSaveSegmentBranch(t, New(""))
}

func TestGetSegmentFound(t *testing.T) {
	adaptertest.TestGetSegmentFound(t, New(""))
}

func TestGetSegmentUpdatedState(t *testing.T) {
	adaptertest.TestGetSegmentUpdatedState(t, New(""))
}

func TestGetSegmentUpdatedMapID(t *testing.T) {
	adaptertest.TestGetSegmentUpdatedMapID(t, New(""))
}

func TestGetSegmentNotFound(t *testing.T) {
	adaptertest.TestGetSegmentNotFound(t, New(""))
}

func TestDeleteSegmentFound(t *testing.T) {
	adaptertest.TestDeleteSegmentFound(t, New(""))
}

func TestDeleteSegmentNotFound(t *testing.T) {
	adaptertest.TestDeleteSegmentNotFound(t, New(""))
}

func TestFindSegmentsAll(t *testing.T) {
	adaptertest.TestFindSegmentsAll(t, New(""))
}

func TestFindSegmentsPagination(t *testing.T) {
	adaptertest.TestFindSegmentsPagination(t, New(""))
}

func TestFindSegmentsEmpty(t *testing.T) {
	adaptertest.TestFindSegmentsEmpty(t, New(""))
}

func TestFindSegmentsSingleTag(t *testing.T) {
	adaptertest.TestFindSegmentsSingleTag(t, New(""))
}

func TestFindSegmentsMultipleTags(t *testing.T) {
	adaptertest.TestFindSegmentsMultipleTags(t, New(""))
}

func TestFindSegmentsMapIDFound(t *testing.T) {
	adaptertest.TestFindSegmentsMapIDFound(t, New(""))
}

func TestFindSegmentsMapIDNotFound(t *testing.T) {
	adaptertest.TestFindSegmentsMapIDNotFound(t, New(""))
}

func TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	adaptertest.TestFindSegmentsPrevLinkHashFound(t, New(""))
}

func TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	adaptertest.TestFindSegmentsPrevLinkHashNotFound(t, New(""))
}

func TestGetMapIDsAll(t *testing.T) {
	adaptertest.TestGetMapIDsAll(t, New(""))
}

func TestGetMapIDsPagination(t *testing.T) {
	adaptertest.TestGetMapIDsPagination(t, New(""))
}

func TestGetMapIDsEmpty(t *testing.T) {
	adaptertest.TestGetMapIDsEmpty(t, New(""))
}

func BenchmarkSaveSegmentNew(b *testing.B) {
	adaptertest.BenchmarkSaveSegmentNew(b, New(""))
}

func BenchmarkSaveSegmentNewParallel(b *testing.B) {
	adaptertest.BenchmarkSaveSegmentNewParallel(b, New(""))
}

func BenchmarkSaveSegmentUpdateState(b *testing.B) {
	adaptertest.BenchmarkSaveSegmentUpdateState(b, New(""))
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	adaptertest.BenchmarkSaveSegmentUpdateStateParallel(b, New(""))
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	adaptertest.BenchmarkSaveSegmentUpdateMapID(b, New(""))
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	adaptertest.BenchmarkSaveSegmentUpdateMapIDParallel(b, New(""))
}

func BenchmarkGetSegmentFound(b *testing.B) {
	adaptertest.BenchmarkGetSegmentFound(b, New(""))
}

func BenchmarkGetSegmentFoundParallel(b *testing.B) {
	adaptertest.BenchmarkGetSegmentFoundParallel(b, New(""))
}

func BenchmarkDeleteSegmentFound(b *testing.B) {
	adaptertest.BenchmarkDeleteSegmentFound(b, New(""))
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
	adaptertest.BenchmarkDeleteSegmentFoundParallel(b, New(""))
}
