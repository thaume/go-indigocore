package adaptertest

import (
	"reflect"
	"sync/atomic"
	"testing"

	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you get an existing segment.
func TestGetSegmentFound(t *testing.T, adapter Adapter) {
	segment1 := RandomSegment()
	linkHash := segment1.Meta["linkHash"].(string)

	adapter.SaveSegment(segment1)

	segment2, err := adapter.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if segment2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(segment1, segment2) {
		t.Fatal("expected segments to be equal")
	}
}

// Tests what happens when you get a segment whose state was updated.
func TestGetSegmentUpdatedState(t *testing.T, adapter Adapter) {
	segment1 := RandomSegment()
	linkHash := segment1.Meta["linkHash"].(string)

	adapter.SaveSegment(segment1)
	segment1 = ChangeSegmentState(segment1)
	adapter.SaveSegment(segment1)

	segment2, err := adapter.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if segment2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(segment1, segment2) {
		t.Fatal("expected segments to be equal")
	}
}

// Tests what happens when you get a segment whose map ID was updated.
func TestGetSegmentUpdatedMapID(t *testing.T, adapter Adapter) {
	segment1 := RandomSegment()
	linkHash := segment1.Meta["linkHash"].(string)

	adapter.SaveSegment(segment1)
	segment1 = ChangeSegmentMapID(segment1)
	adapter.SaveSegment(segment1)

	segment2, err := adapter.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if segment2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(segment1, segment2) {
		t.Fatal("expected segments to be equal")
	}
}

// Tests what happens when you get a nonexistent segment.
func TestGetSegmentNotFound(t *testing.T, adapter Adapter) {
	segment, err := adapter.GetSegment(RandomString(32))

	if err != nil {
		t.Fatal(err)
	}

	if segment != nil {
		t.Fatal("expected segment to be nil")
	}
}

func BenchmarkGetSegmentFound(b *testing.B, adapter Adapter) {
	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		linkHashes[i] = segment.Meta["linkHash"].(string)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if segment, err := adapter.GetSegment(linkHashes[i]); err != nil {
			b.Fatal(err)
		} else if segment == nil {
			b.Fatal("expected segment")
		}
	}
}

func BenchmarkGetSegmentFoundParallel(b *testing.B, adapter Adapter) {
	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		linkHashes[i] = segment.Meta["linkHash"].(string)
	}

	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if segment, err := adapter.GetSegment(linkHashes[i]); err != nil {
				b.Fatal(err)
			} else if segment == nil {
				b.Fatal("expected segment")
			}
		}
	})
}
