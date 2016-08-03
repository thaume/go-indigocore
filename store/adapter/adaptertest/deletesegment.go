package adaptertest

import (
	"reflect"
	"sync/atomic"
	"testing"

	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you delete an existing segments.
func TestDeleteSegmentFound(t *testing.T, a Adapter) {
	s1 := RandomSegment()
	a.SaveSegment(s1)

	linkHash := s1.Meta["linkHash"].(string)

	s2, err := a.DeleteSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if s2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(s1, s2) {
		t.Fatal("expected segments to be equal")
	}

	s2, err = a.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if s2 != nil {
		t.Fatal("expected segment to be nil")
	}
}

// Tests what happens when you delete a nonexistent segment.
func TestDeleteSegmentNotFound(t *testing.T, a Adapter) {
	s, err := a.DeleteSegment(RandomString(32))

	if err != nil {
		t.Fatal(err)
	}

	if s != nil {
		t.Fatal("expected segment to be nil")
	}
}

func BenchmarkDeleteSegmentFound(b *testing.B, a Adapter) {
	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		s := RandomSegment()
		a.SaveSegment(s)
		linkHashes[i] = s.Meta["linkHash"].(string)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if s, err := a.DeleteSegment(linkHashes[i]); err != nil {
			b.Fatal(err)
		} else if s == nil {
			b.Fatal("expected segment")
		}
	}
}

func BenchmarkDeleteSegmentFoundParallel(b *testing.B, a Adapter) {
	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		s := RandomSegment()
		a.SaveSegment(s)
		linkHashes[i] = s.Meta["linkHash"].(string)
	}

	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if s, err := a.DeleteSegment(linkHashes[i]); err != nil {
				b.Fatal(err)
			} else if s == nil {
				b.Fatal("expected segment")
			}
		}
	})
}
