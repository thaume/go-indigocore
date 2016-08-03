package storetesting

import (
	"sync/atomic"
	"testing"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/store"
)

// TestSaveSegmentNew tests what happens when you save a new segment.
func TestSaveSegmentNew(t *testing.T, a store.Adapter) {
	s := cstesting.RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// TestSaveSegmentUpdateState tests what happens when you update the state of a segment.
func TestSaveSegmentUpdateState(t *testing.T, a store.Adapter) {
	s := cstesting.RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}

	cstesting.ChangeSegmentState(s)

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// TestSaveSegmentUpdateMapID tests what happens when you update the map ID of a segment.
func TestSaveSegmentUpdateMapID(t *testing.T, a store.Adapter) {
	s1 := cstesting.RandomSegment()

	if err := a.SaveSegment(s1); err != nil {
		t.Fatal(err)
	}

	s2 := cstesting.ChangeSegmentMapID(s1)

	if err := a.SaveSegment(s2); err != nil {
		t.Fatal(err)
	}
}

// TestSaveSegmentBranch tests what happens when you save a segment with a previous link hash.
func TestSaveSegmentBranch(t *testing.T, a store.Adapter) {
	s := cstesting.RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}

	s = cstesting.RandomBranch(s)

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// BenchmarkSaveSegmentNew benchmarks saving new segments.
func BenchmarkSaveSegmentNew(b *testing.B, a store.Adapter) {
	slice := make([]*cs.Segment, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = cstesting.RandomSegment()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSaveSegmentNewParallel benchmarks saving new segments in parallel.
func BenchmarkSaveSegmentNewParallel(b *testing.B, a store.Adapter) {
	slice := make([]*cs.Segment, b.N)

	for i := 0; i < b.N; i++ {
		slice[i] = cstesting.RandomSegment()
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

// BenchmarkSaveSegmentUpdateState benchmarks updating segments states.
func BenchmarkSaveSegmentUpdateState(b *testing.B, a store.Adapter) {
	slice := make([]*cs.Segment, b.N)

	for i := 0; i < b.N; i++ {
		s := cstesting.RandomSegment()
		a.SaveSegment(s)
		slice[i] = cstesting.ChangeSegmentState(s)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSaveSegmentUpdateStateParallel benchmarks updating segments states in parallel.
func BenchmarkSaveSegmentUpdateStateParallel(b *testing.B, a store.Adapter) {
	slice := make([]*cs.Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := cstesting.RandomSegment()
		a.SaveSegment(segment)
		slice[i] = cstesting.ChangeSegmentState(segment)
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

// BenchmarkSaveSegmentUpdateMapID benchmarks updating segment map IDs.
func BenchmarkSaveSegmentUpdateMapID(b *testing.B, a store.Adapter) {
	slice := make([]*cs.Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := cstesting.RandomSegment()
		a.SaveSegment(segment)
		slice[i] = cstesting.ChangeSegmentMapID(segment)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSaveSegmentUpdateMapIDParallel benchmarks updating segment map IDs in parallel.
func BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B, a store.Adapter) {
	slice := make([]*cs.Segment, b.N)

	for i := 0; i < b.N; i++ {
		segment := cstesting.RandomSegment()
		a.SaveSegment(segment)
		slice[i] = cstesting.ChangeSegmentMapID(segment)
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
