package adaptertest

import (
	"sync/atomic"
	"testing"

	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you save a new segment.
func TestSaveSegmentNew(t *testing.T, a Adapter) {
	s := RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// Tests what happens when you update the state of a segment.
func TestSaveSegmentUpdateState(t *testing.T, a Adapter) {
	s := RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}

	ChangeSegmentState(s)

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// Tests what happens when you update the map ID of a segment.
func TestSaveSegmentUpdateMapID(t *testing.T, a Adapter) {
	s1 := RandomSegment()

	if err := a.SaveSegment(s1); err != nil {
		t.Fatal(err)
	}

	s2 := ChangeSegmentMapID(s1)

	if err := a.SaveSegment(s2); err != nil {
		t.Fatal(err)
	}
}

// Tests what happens when you save a segment with a previous link hash.
func TestSaveSegmentBranch(t *testing.T, a Adapter) {
	s := RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}

	s = RandomBranch(s)

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkSaveSegmentNew(b *testing.B, a Adapter) {
	slice := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = RandomSegment()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSaveSegmentNewParallel(b *testing.B, a Adapter) {
	slice := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = RandomSegment()
	}

	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if err := a.SaveSegment(slice[i]); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSaveSegmentUpdateState(b *testing.B, a Adapter) {
	slice := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		s := RandomSegment()
		a.SaveSegment(s)
		slice[i] = ChangeSegmentState(s)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B, a Adapter) {
	slice := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		a.SaveSegment(segment)
		slice[i] = ChangeSegmentState(segment)
	}
	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if err := a.SaveSegment(slice[i]); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B, a Adapter) {
	slice := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		a.SaveSegment(segment)
		slice[i] = ChangeSegmentMapID(segment)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B, a Adapter) {
	slice := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		a.SaveSegment(segment)
		slice[i] = ChangeSegmentMapID(segment)
	}
	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if err := a.SaveSegment(slice[i]); err != nil {
				b.Fatal(err)
			}
		}
	})
}
