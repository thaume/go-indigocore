// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/store"
	"github.com/stratumn/go/testutil"
	"github.com/stratumn/go/types"
)

// TestFindSegments tests what happens when you search for all segments.
func (f Factory) TestFindSegments(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < 100; i++ {
		a.SaveSegment(cstesting.RandomSegment())
	}

	slice, err := a.FindSegments(&store.Filter{})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 100; got != want {
		t.Errorf("len(slice) = %d want %d", got, want)
	}

	wantLTE := 100.0
	for _, s := range slice {
		got := s.Link.Meta["priority"].(float64)
		if got > wantLTE {
			t.Errorf("priority = %f want <= %f", got, wantLTE)
		}
		wantLTE = got
	}
}

// TestFindSegmentsPagination tests what happens when you search with pagination.
func (f Factory) TestFindSegmentsPagination(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < 100; i++ {
		a.SaveSegment(cstesting.RandomSegment())
	}

	limit := 10 + rand.Intn(10)
	slice, err := a.FindSegments(&store.Filter{
		Pagination: store.Pagination{
			Offset: rand.Intn(40),
			Limit:  limit,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), limit; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}

	wantLTE := 100.0
	for _, s := range slice {
		got := s.Link.Meta["priority"].(float64)
		if got > wantLTE {
			t.Errorf("priority = %f want <= %f", got, wantLTE)
		}
		wantLTE = got
	}
}

// TestFindSegmentEmpty tests what happens when there are no matches.
func (f Factory) TestFindSegmentEmpty(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < 100; i++ {
		a.SaveSegment(cstesting.RandomSegment())
	}

	slice, err := a.FindSegments(&store.Filter{
		Tags: []string{"blablabla"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 0; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestFindSegmentsSingleTag tests what happens when you search with only one tag.
func (f Factory) TestFindSegmentsSingleTag(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	tag1 := testutil.RandomString(5)
	tag2 := testutil.RandomString(5)

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, testutil.RandomString(5)}
		a.SaveSegment(s)
	}

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, tag2, testutil.RandomString(5)}
		a.SaveSegment(s)
	}

	slice, err := a.FindSegments(&store.Filter{
		Tags: []string{tag1},
	})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 20; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestFindSegmentsMultipleTags tests what happens when you search with more than one tag.
func (f Factory) TestFindSegmentsMultipleTags(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	tag1 := testutil.RandomString(5)
	tag2 := testutil.RandomString(5)

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, testutil.RandomString(5)}
		a.SaveSegment(s)
	}

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, tag2, testutil.RandomString(5)}
		a.SaveSegment(s)
	}

	slice, err := a.FindSegments(&store.Filter{
		Tags: []string{tag2, tag1},
	})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 10; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestFindSegmentsMapID tests whan happens when you search for an existing map ID.
func (f Factory) TestFindSegmentsMapID(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			s := cstesting.RandomSegment()
			s.Link.Meta["mapId"] = fmt.Sprintf("map%d", i)
			a.SaveSegment(s)
		}
	}

	slice, err := a.FindSegments(&store.Filter{
		MapID: "map1",
	})
	if err != nil {
		t.Fatal(err)
	}

	if got := slice; got == nil {
		t.Fatal("slice = nit want cs.SegmentSlice")
	}
	if got, want := len(slice), 10; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestFindSegmentsMapIDNotFound tests whan happens when you search for a nonexistent map ID.
func (f Factory) TestFindSegmentsMapIDNotFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	slice, err := a.FindSegments(&store.Filter{
		MapID: testutil.RandomString(10),
	})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 0; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestFindSegmentsPrevLinkHash tests whan happens when you search for an existing previous link hash.
func (f Factory) TestFindSegmentsPrevLinkHash(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	s := cstesting.RandomSegment()
	a.SaveSegment(s)

	for i := 0; i < 10; i++ {
		a.SaveSegment(cstesting.RandomBranch(s))
	}

	linkHash, err := types.NewBytes32FromString(s.Meta["linkHash"].(string))
	if err != nil {
		t.Fatal(err)
	}

	slice, err := a.FindSegments(&store.Filter{
		PrevLinkHash: linkHash,
	})
	if err != nil {
		t.Fatal(err)
	}

	if got := slice; got == nil {
		t.Fatal("slice = nit want cs.SegmentSlice")
	}
	if got, want := len(slice), 10; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestFindSegmentsPrevLinkHashNotFound tests whan happens when you search for a nonexistent previous link hash.
func (f Factory) TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	slice, err := a.FindSegments(&store.Filter{
		PrevLinkHash: testutil.RandomHash(),
	})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 0; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}
