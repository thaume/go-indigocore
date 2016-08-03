package adaptertest

import (
	"reflect"
	"sync/atomic"
	"testing"

	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you delete an existing segments.
func TestDeleteSegmentFound(t *testing.T, adapter Adapter) {
	segment1 := RandomSegment()
	adapter.SaveSegment(segment1)

	linkHash := segment1.Meta["linkHash"].(string)

	segment2, err := adapter.DeleteSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if segment2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(segment1, segment2) {
		t.Fatal("expected segments to be equal")
	}

	segment2, err = adapter.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if segment2 != nil {
		t.Fatal("expected segment to be nil")
	}
}

// Tests what happens when you delete a nonexistent segment.
func TestDeleteSegmentNotFound(t *testing.T, adapter Adapter) {
	segment, err := adapter.DeleteSegment(RandomString(32))

	if err != nil {
		t.Fatal(err)
	}

	if segment != nil {
		t.Fatal("expected segment to be nil")
	}
}

func BenchmarkDeleteSegmentFound(b *testing.B, adapter Adapter) {
	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		linkHashes[i] = segment.Meta["linkHash"].(string)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if segment, err := adapter.DeleteSegment(linkHashes[i]); err != nil {
			b.Fatal(err)
		} else if segment == nil {
			b.Fatal("expected segment")
		}
	}
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B, adapter Adapter) {
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

			if segment, err := adapter.DeleteSegment(linkHashes[i]); err != nil {
				b.Fatal(err)
			} else if segment == nil {
				b.Fatal("expected segment")
			}
		}
	})
}
