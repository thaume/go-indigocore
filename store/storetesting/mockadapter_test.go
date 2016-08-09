// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetesting

import (
	"reflect"
	"testing"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/store"
)

func TestMockAdapter_GetInfo(t *testing.T) {
	a := &MockAdapter{}

	_, err := a.GetInfo()

	if err != nil {
		t.Fatal("unexpected error")
	}

	a.MockGetInfo.Fn = func() (interface{}, error) { return map[string]string{"name": "test"}, nil }
	info, err := a.GetInfo()

	if err != nil {
		t.Fatal("unexpected error")
	}

	if info.(map[string]string)["name"] != "test" {
		t.Fatal("unexpect info")
	}

	if a.MockGetInfo.CalledCount != 2 {
		t.Fatal("unexpected MockGetInfo.CalledCount value")
	}
}

func TestMockAdapter_SaveSegment(t *testing.T) {
	a := &MockAdapter{}
	s := cstesting.RandomSegment()

	err := a.SaveSegment(s)

	if err != nil {
		t.Fatal("unexpected error")
	}

	a.MockSaveSegment.Fn = func(s *cs.Segment) error { return nil }
	err = a.SaveSegment(s)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if a.MockSaveSegment.CalledCount != 2 {
		t.Fatal("unexpected MockSaveSegment.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockSaveSegment.CalledWith, []*cs.Segment{s, s}) {
		t.Fatal("unexpected MockSaveSegment.LastCalledWith value")
	}

	if a.MockSaveSegment.LastCalledWith != s {
		t.Fatal("unexpected MockSaveSegment.LastCalledWith value")
	}
}

func TestMockAdapter_GetSegment(t *testing.T) {
	a := &MockAdapter{}

	_, err := a.GetSegment("abcdef")

	if err != nil {
		t.Fatal("unexpected error")
	}

	s1 := cstesting.RandomSegment()
	a.MockGetSegment.Fn = func(linkHash string) (*cs.Segment, error) { return s1, nil }
	s2, err := a.GetSegment("ghij")

	if err != nil {
		t.Fatal("unexpected error")
	}

	if s1 != s2 {
		t.Fatal("expected segments to be equal")
	}

	if a.MockGetSegment.CalledCount != 2 {
		t.Fatal("unexpected MockGetSegment.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockGetSegment.CalledWith, []string{"abcdef", "ghij"}) {
		t.Fatal("unexpected MockGetSegment.LastCalledWith value")
	}

	if a.MockGetSegment.LastCalledWith != "ghij" {
		t.Fatal("unexpected MockGetSegment.LastCalledWith value")
	}
}

func TestMockAdapter_DeleteSegment(t *testing.T) {
	a := &MockAdapter{}

	_, err := a.DeleteSegment("abcdef")

	if err != nil {
		t.Fatal("unexpected error")
	}

	s1 := cstesting.RandomSegment()
	a.MockDeleteSegment.Fn = func(linkHash string) (*cs.Segment, error) { return s1, nil }
	s2, err := a.DeleteSegment("ghij")

	if err != nil {
		t.Fatal("unexpected error")
	}

	if s1 != s2 {
		t.Fatal("expected segments to be equal")
	}

	if a.MockDeleteSegment.CalledCount != 2 {
		t.Fatal("unexpected MockDeleteSegment.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockDeleteSegment.CalledWith, []string{"abcdef", "ghij"}) {
		t.Fatal("unexpected MockDeleteSegment.LastCalledWith value")
	}

	if a.MockDeleteSegment.LastCalledWith != "ghij" {
		t.Fatal("unexpected MockDeleteSegment.LastCalledWith value")
	}
}

func TestMockAdapter_FindSegments(t *testing.T) {
	a := &MockAdapter{}

	_, err := a.FindSegments(nil)

	if err != nil {
		t.Fatal("unexpected error")
	}

	s := cstesting.RandomSegment()
	a.MockFindSegments.Fn = func(*store.Filter) (cs.SegmentSlice, error) { return cs.SegmentSlice{s}, nil }
	f := store.Filter{PrevLinkHash: "test"}
	slice, err := a.FindSegments(&f)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if !reflect.DeepEqual(slice, cs.SegmentSlice{s}) {
		t.Fatal("expected segment slices to be equal")
	}

	if a.MockFindSegments.CalledCount != 2 {
		t.Fatal("unexpected MockFindSegments.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockFindSegments.CalledWith, []*store.Filter{nil, &f}) {
		t.Fatal("unexpected MockFindSegments.LastCalledWith value")
	}

	if a.MockFindSegments.LastCalledWith != &f {
		t.Fatal("unexpected MockFindSegments.LastCalledWith value")
	}
}

func TestMockAdapter_GetMapIDs(t *testing.T) {
	a := &MockAdapter{}

	_, err := a.GetMapIDs(nil)

	if err != nil {
		t.Fatal("unexpected error")
	}

	a.MockGetMapIDs.Fn = func(*store.Pagination) ([]string, error) { return []string{"one", "two"}, nil }
	p := store.Pagination{Offset: 10}
	slice, err := a.GetMapIDs(&p)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if !reflect.DeepEqual(slice, []string{"one", "two"}) {
		t.Fatal("expected segment slices to be equal")
	}

	if a.MockGetMapIDs.CalledCount != 2 {
		t.Fatal("unexpected MockGetMapIDs.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockGetMapIDs.CalledWith, []*store.Pagination{nil, &p}) {
		t.Fatal("unexpected MockGetMapIDs.LastCalledWith value")
	}

	if a.MockGetMapIDs.LastCalledWith != &p {
		t.Fatal("unexpected MockGetMapIDs.LastCalledWith value")
	}
}
