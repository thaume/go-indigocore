package adaptertest

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you search for all segments.
func TestFindSegmentsAll(t *testing.T, a Adapter) {
	for i := 0; i < 100; i++ {
		a.SaveSegment(RandomSegment())
	}

	slice, err := a.FindSegments(&Filter{})

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

// Tests what happens when you search with pagination.
func TestFindSegmentsPagination(t *testing.T, a Adapter) {
	for i := 0; i < 100; i++ {
		a.SaveSegment(RandomSegment())
	}

	limit := 10 + rand.Intn(10)

	slice, err := a.FindSegments(&Filter{
		Pagination: Pagination{
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

// Tests what happens when there are no matches.
func TestFindSegmentsEmpty(t *testing.T, a Adapter) {
	for i := 0; i < 100; i++ {
		a.SaveSegment(RandomSegment())
	}

	slice, err := a.FindSegments(&Filter{
		Tags: []string{"blablabla"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}

// Tests what happens when you search with only one tag.
func TestFindSegmentsSingleTag(t *testing.T, a Adapter) {
	tag1 := RandomString(5)
	tag2 := RandomString(5)

	for i := 0; i < 10; i++ {
		s := RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, RandomString(5)}
		a.SaveSegment(s)
	}

	for i := 0; i < 10; i++ {
		s := RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, tag2, RandomString(5)}
		a.SaveSegment(s)
	}

	slice, err := a.FindSegments(&Filter{
		Tags: []string{tag1},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 20 {
		t.Fatalf("expected segments length to be 20")
	}
}

// Tests what happens when you search with more than one tag.
func TestFindSegmentsMultipleTags(t *testing.T, a Adapter) {
	tag1 := RandomString(5)
	tag2 := RandomString(5)

	for i := 0; i < 10; i++ {
		s := RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, RandomString(5)}
		a.SaveSegment(s)
	}

	for i := 0; i < 10; i++ {
		s := RandomSegment()
		s.Link.Meta["tags"] = []interface{}{tag1, tag2, RandomString(5)}
		a.SaveSegment(s)
	}

	slice, err := a.FindSegments(&Filter{
		Tags: []string{tag2, tag1},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 10 {
		t.Fatalf("expected segments length to be 10")
	}
}

// Tests whan happens when you search for an existing map ID.
func TestFindSegmentsMapIDFound(t *testing.T, a Adapter) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			s := RandomSegment()
			s.Link.Meta["mapId"] = fmt.Sprintf("map%d", i)
			a.SaveSegment(s)
		}
	}

	slice, err := a.FindSegments(&Filter{
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

// Tests whan happens when you search for a nonexistent map ID.
func TestFindSegmentsMapIDNotFound(t *testing.T, a Adapter) {
	slice, err := a.FindSegments(&Filter{
		MapID: RandomString(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}

// Tests whan happens when you search for an existing previous link hash.
func TestFindSegmentsPrevLinkHashFound(t *testing.T, a Adapter) {
	s := RandomSegment()
	a.SaveSegment(s)

	for i := 0; i < 10; i++ {
		a.SaveSegment(RandomBranch(s))
	}

	slice, err := a.FindSegments(&Filter{
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

// Tests whan happens when you search for a nonexistent previous link hash.
func TestFindSegmentsPrevLinkHashNotFound(t *testing.T, a Adapter) {
	slice, err := a.FindSegments(&Filter{
		PrevLinkHash: RandomString(32),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}
