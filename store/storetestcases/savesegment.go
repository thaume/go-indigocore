// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"sync/atomic"
	"testing"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/cs/cstesting"
)

// TestSaveSegmentNew tests what happens when you save a new segment.
func (f Factory) TestSaveSegmentNew(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s := cstesting.RandomSegment()

	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// TestSaveSegmentUpdateState tests what happens when you update the state of a segment.
func (f Factory) TestSaveSegmentUpdateState(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) TestSaveSegmentUpdateMapID(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) TestSaveSegmentBranch(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) BenchmarkSaveSegmentNew(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) BenchmarkSaveSegmentNewParallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) BenchmarkSaveSegmentUpdateState(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) BenchmarkSaveSegmentUpdateStateParallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) BenchmarkSaveSegmentUpdateMapID(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
func (f Factory) BenchmarkSaveSegmentUpdateMapIDParallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

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
