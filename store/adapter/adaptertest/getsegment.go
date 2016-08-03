package adaptertest

import (
	"reflect"
	"sync/atomic"
	"testing"

	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you get an existing segment.
func TestGetSegmentFound(t *testing.T, a Adapter) {
	s1 := RandomSegment()
	linkHash := s1.Meta["linkHash"].(string)

	a.SaveSegment(s1)

	s2, err := a.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if s2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(s1, s2) {
		t.Fatal("expected segments to be equal")
	}
}

// Tests what happens when you get a segment whose state was updated.
func TestGetSegmentUpdatedState(t *testing.T, a Adapter) {
	s1 := RandomSegment()
	linkHash := s1.Meta["linkHash"].(string)

	a.SaveSegment(s1)
	s1 = ChangeSegmentState(s1)
	a.SaveSegment(s1)

	s2, err := a.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if s2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(s1, s2) {
		t.Fatal("expected segments to be equal")
	}
}

// Tests what happens when you get a segment whose map ID was updated.
func TestGetSegmentUpdatedMapID(t *testing.T, a Adapter) {
	s1 := RandomSegment()
	linkHash := s1.Meta["linkHash"].(string)

	a.SaveSegment(s1)
	s1 = ChangeSegmentMapID(s1)
	a.SaveSegment(s1)

	s2, err := a.GetSegment(linkHash)

	if err != nil {
		t.Fatal(err)
	}

	if s2 == nil {
		t.Fatal("expected segment not to be nil")
	}

	if !reflect.DeepEqual(s1, s2) {
		t.Fatal("expected segments to be equal")
	}
}

// Tests what happens when you get a nonexistent segment.
func TestGetSegmentNotFound(t *testing.T, a Adapter) {
	s, err := a.GetSegment(RandomString(32))

	if err != nil {
		t.Fatal(err)
	}

	if s != nil {
		t.Fatal("expected segment to be nil")
	}
}

func BenchmarkGetSegmentFound(b *testing.B, a Adapter) {
	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		s := RandomSegment()
		a.SaveSegment(s)
		linkHashes[i] = s.Meta["linkHash"].(string)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if s, err := a.GetSegment(linkHashes[i]); err != nil {
			b.Fatal(err)
		} else if s == nil {
			b.Fatal("expected segment")
		}
	}
}

func BenchmarkGetSegmentFoundParallel(b *testing.B, a Adapter) {
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

			if s, err := a.GetSegment(linkHashes[i]); err != nil {
				b.Fatal(err)
			} else if s == nil {
				b.Fatal("expected segment")
			}
		}
	})
}
