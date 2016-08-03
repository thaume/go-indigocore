package adaptertest

import (
	"sync/atomic"
	"testing"

	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you save a new segment.
func TestSaveSegmentNew(t *testing.T, adapter Adapter) {
	segment := RandomSegment()

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}
}

// Tests what happens when you update the state of a segment.
func TestSaveSegmentUpdateState(t *testing.T, adapter Adapter) {
	segment := RandomSegment()

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}

	ChangeSegmentState(segment)

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}
}

// Tests what happens when you update the map ID of a segment.
func TestSaveSegmentUpdateMapID(t *testing.T, adapter Adapter) {
	segment := RandomSegment()

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}

	ChangeSegmentMapID(segment)

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}
}

// Tests what happens when you save a segment with a previous link hash.
func TestSaveSegmentBranch(t *testing.T, adapter Adapter) {
	segment := RandomSegment()

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}

	segment = RandomBranch(segment)

	if err := adapter.SaveSegment(segment); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkSaveSegmentNew(b *testing.B, adapter Adapter) {
	segments := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segments[i] = RandomSegment()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := adapter.SaveSegment(segments[i]); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSaveSegmentNewParallel(b *testing.B, adapter Adapter) {
	segments := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segments[i] = RandomSegment()
	}

	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if err := adapter.SaveSegment(segments[i]); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSaveSegmentUpdateState(b *testing.B, adapter Adapter) {
	segments := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		segments[i] = ChangeSegmentState(segment)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := adapter.SaveSegment(segments[i]); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B, adapter Adapter) {
	segments := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		segments[i] = ChangeSegmentState(segment)
	}
	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if err := adapter.SaveSegment(segments[i]); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSaveSegmentUpdateMapID(b *testing.B, adapter Adapter) {
	segments := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		segments[i] = ChangeSegmentMapID(segment)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := adapter.SaveSegment(segments[i]); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B, adapter Adapter) {
	segments := make([]*Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := RandomSegment()
		adapter.SaveSegment(segment)
		segments[i] = ChangeSegmentMapID(segment)
	}
	var counter uint64

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1

			if err := adapter.SaveSegment(segments[i]); err != nil {
				b.Fatal(err)
			}
		}
	})
}
