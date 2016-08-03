package adaptertest

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you search for all segments.
func TestFindSegmentsAll(t *testing.T, adapter Adapter) {
	for i := 0; i < 100; i++ {
		adapter.SaveSegment(RandomSegment())
	}

	segments, err := adapter.FindSegments(&Filter{})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 100 {
		t.Fatal("expected segments length to be 100")
	}

	lastPriority := 100.0

	for _, segment := range segments {
		priority := segment.Link.Meta["priority"].(float64)

		if priority > lastPriority {
			t.Fatal("segments not ordered by priority")
		}

		lastPriority = priority
	}
}

// Tests what happens when you search with pagination.
func TestFindSegmentsPagination(t *testing.T, adapter Adapter) {
	for i := 0; i < 100; i++ {
		adapter.SaveSegment(RandomSegment())
	}

	limit := 10 + rand.Intn(10)

	segments, err := adapter.FindSegments(&Filter{
		Pagination: Pagination{
			Offset: rand.Intn(40),
			Limit:  limit,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != limit {
		t.Fatalf("expected segments length to be %d", limit)
	}

	lastPriority := 100.0

	for _, segment := range segments {
		priority := segment.Link.Meta["priority"].(float64)

		if priority > lastPriority {
			t.Fatal("segments not ordered by priority")
		}

		lastPriority = priority
	}
}

// Tests what happens when there are no matches.
func TestFindSegmentsEmpty(t *testing.T, adapter Adapter) {
	for i := 0; i < 100; i++ {
		adapter.SaveSegment(RandomSegment())
	}

	segments, err := adapter.FindSegments(&Filter{
		Tags: []string{"blablabla"},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}

// Tests what happens when you search with only one tag.
func TestFindSegmentsSingleTag(t *testing.T, adapter Adapter) {
	tag1 := RandomString(5)
	tag2 := RandomString(5)

	for i := 0; i < 10; i++ {
		segment := RandomSegment()
		segment.Link.Meta["tags"] = []interface{}{tag1, RandomString(5)}
		adapter.SaveSegment(segment)
	}

	for i := 0; i < 10; i++ {
		segment := RandomSegment()
		segment.Link.Meta["tags"] = []interface{}{tag1, tag2, RandomString(5)}
		adapter.SaveSegment(segment)
	}

	segments, err := adapter.FindSegments(&Filter{
		Tags: []string{tag1},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 20 {
		t.Fatalf("expected segments length to be 20")
	}
}

// Tests what happens when you search with more than one tag.
func TestFindSegmentsMultipleTags(t *testing.T, adapter Adapter) {
	tag1 := RandomString(5)
	tag2 := RandomString(5)

	for i := 0; i < 10; i++ {
		segment := RandomSegment()
		segment.Link.Meta["tags"] = []interface{}{tag1, RandomString(5)}
		adapter.SaveSegment(segment)
	}

	for i := 0; i < 10; i++ {
		segment := RandomSegment()
		segment.Link.Meta["tags"] = []interface{}{tag1, tag2, RandomString(5)}
		adapter.SaveSegment(segment)
	}

	segments, err := adapter.FindSegments(&Filter{
		Tags: []string{tag2, tag1},
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 10 {
		t.Fatalf("expected segments length to be 10")
	}
}

// Tests whan happens when you search for an existing map ID.
func TestFindSegmentsMapIDFound(t *testing.T, adapter Adapter) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			segment := RandomSegment()
			segment.Link.Meta["mapId"] = fmt.Sprintf("map%d", i)
			adapter.SaveSegment(segment)
		}
	}

	segments, err := adapter.FindSegments(&Filter{
		MapID: "map1",
	})

	if err != nil {
		t.Fatal(err)
	}

	if segments == nil {
		t.Fatal("expected segments not to be nil")
	}

	if len(segments) != 10 {
		t.Fatal("expected segments length to be 10")
	}
}

// Tests whan happens when you search for a nonexistent map ID.
func TestFindSegmentsMapIDNotFound(t *testing.T, adapter Adapter) {
	segments, err := adapter.FindSegments(&Filter{
		MapID: RandomString(10),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}

// Tests whan happens when you search for an existing previous link hash.
func TestFindSegmentsPrevLinkHashFound(t *testing.T, adapter Adapter) {
	segment := RandomSegment()
	adapter.SaveSegment(segment)

	for i := 0; i < 10; i++ {
		adapter.SaveSegment(RandomBranch(segment))
	}

	segments, err := adapter.FindSegments(&Filter{
		PrevLinkHash: segment.Meta["linkHash"].(string),
	})

	if err != nil {
		t.Fatal(err)
	}

	if segments == nil {
		t.Fatal("expected segments not to be nil")
	}

	if len(segments) != 10 {
		t.Fatal("expected segments length to be 10")
	}
}

// Tests whan happens when you search for a nonexistent previous link hash.
func TestFindSegmentsPrevLinkHashNotFound(t *testing.T, adapter Adapter) {
	segments, err := adapter.FindSegments(&Filter{
		PrevLinkHash: RandomString(32),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(segments) != 0 {
		t.Fatal("expected segments length to be 0")
	}
}
