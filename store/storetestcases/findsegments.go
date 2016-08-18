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
)

// TestFindSegmentsAll tests what happens when you search for all segments.
func (f Factory) TestFindSegmentsAll(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	for i := 0; i < 100; i++ {
		a.SaveSegment(cstesting.RandomSegment())
	}

	slice, err := a.FindSegments(&store.Filter{})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 100 {
		t.Fatal("expected segments length to be 100")
	}

	lastPriority := 100.0

	for _, s := range slice {
		priority := s.Link.Meta["priority"].(float64)

		if priority > lastPriority {
			t.Fatal("segments not ordered by priority")
		}

		lastPriority = priority
	}
}

// TestFindSegmentsPagination tests what happens when you search with pagination.
func (f Factory) TestFindSegmentsPagination(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
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

	if len(slice) != limit {
		t.Fatalf("expected segments length to be %d", limit)
	}

	lastPriority := 100.0

	for _, s := range slice {
		priority := s.Link.Meta["priority"].(float64)

		if priority > lastPriority {
			t.Fatal("segments not ordered by priority")
		}

		lastPriority = priority
	}
}

// TestFindSegmentsEmpty tests what happens when there are no matches.
func (f Factory) TestFindSegmentsEmpty(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
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

	if len(slice) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}

// TestFindSegmentsSingleTag tests what happens when you search with only one tag.
func (f Factory) TestFindSegmentsSingleTag(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	tag1 := cstesting.RandomString(5)
	tag2 := cstesting.RandomString(5)

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, cstesting.RandomString(5)}
		a.SaveSegment(s)
	}

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, tag2, cstesting.RandomString(5)}
		a.SaveSegment(s)
	}

	slice, err := a.FindSegments(&store.Filter{
		Tags: []string{tag1},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 20 {
		t.Fatalf("expected segments length to be 20")
	}
}

// TestFindSegmentsMultipleTags tests what happens when you search with more than one tag.
func (f Factory) TestFindSegmentsMultipleTags(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	tag1 := cstesting.RandomString(5)
	tag2 := cstesting.RandomString(5)

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, cstesting.RandomString(5)}
		a.SaveSegment(s)
	}

	for i := 0; i < 10; i++ {
		s := cstesting.RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, tag2, cstesting.RandomString(5)}
		a.SaveSegment(s)
	}

	slice, err := a.FindSegments(&store.Filter{
		Tags: []string{tag2, tag1},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 10 {
		t.Fatalf("expected segments length to be 10")
	}
}

// TestFindSegmentsMapIDFound tests whan happens when you search for an existing map ID.
func (f Factory) TestFindSegmentsMapIDFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
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

	if slice == nil {
		t.Fatal("expected segments not to be nil")
	}

	if len(slice) != 10 {
		t.Fatal("expected segments length to be 10")
	}
}

// TestFindSegmentsMapIDNotFound tests whan happens when you search for a nonexistent map ID.
func (f Factory) TestFindSegmentsMapIDNotFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	slice, err := a.FindSegments(&store.Filter{
		MapID: cstesting.RandomString(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}

// TestFindSegmentsPrevLinkHashFound tests whan happens when you search for an existing previous link hash.
func (f Factory) TestFindSegmentsPrevLinkHashFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s := cstesting.RandomSegment()
	a.SaveSegment(s)

	for i := 0; i < 10; i++ {
		a.SaveSegment(cstesting.RandomBranch(s))
	}

	slice, err := a.FindSegments(&store.Filter{
		PrevLinkHash: s.Meta["linkHash"].(string),
	})

	if err != nil {
		t.Fatal(err)
	}

	if slice == nil {
		t.Fatal("expected segments not to be nil")
	}

	if len(slice) != 10 {
		t.Fatal("expected segments length to be 10")
	}
}

// TestFindSegmentsPrevLinkHashNotFound tests whan happens when you search for a nonexistent previous link hash.
func (f Factory) TestFindSegmentsPrevLinkHashNotFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	slice, err := a.FindSegments(&store.Filter{
		PrevLinkHash: cstesting.RandomString(32),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}
