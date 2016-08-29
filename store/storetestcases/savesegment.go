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

// TestSaveSegment tests what happens when you save a new segment.
func (f Factory) TestSaveSegment(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	s := cstesting.RandomSegment()
	if err := a.SaveSegment(s); err != nil {
		t.Fatal(err)
	}
}

// TestSaveSegment_updatedState tests what happens when you update the state of a segment.
func (f Factory) TestSaveSegment_updatedState(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
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

// TestSaveSegment_updatedMapID tests what happens when you update the map ID of a segment.
func (f Factory) TestSaveSegment_updatedMapID(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
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

// TestSaveSegment_branch tests what happens when you save a segment with a previous link hash.
func (f Factory) TestSaveSegment_branch(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
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

// BenchmarkSaveSegment benchmarks saving new segments.
func (f Factory) BenchmarkSaveSegment(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	slice := make([]*cs.Segment, b.N)
	for i := 0; i < b.N; i++ {
		slice[i] = cstesting.RandomSegment()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Error(err)
		}
	}
}

// BenchmarkSaveSegment_parallel benchmarks saving new segments in parallel.
func (f Factory) BenchmarkSaveSegment_parallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
				b.Error(err)
			}
		}
	})
}

// BenchmarkSaveSegment_updatedState benchmarks updating segments states.
func (f Factory) BenchmarkSaveSegment_updatedState(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
			b.Error(err)
		}
	}
}

// BenchmarkSaveSegment_updatedStateParallel benchmarks updating segments states in parallel.
func (f Factory) BenchmarkSaveSegment_updatedStateParallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
				b.Error(err)
			}
		}
	})
}

// BenchmarkSaveSegment_updatedMapID benchmarks updating segment map IDs.
func (f Factory) BenchmarkSaveSegment_updatedMapID(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
			b.Error(err)
		}
	}
}

// BenchmarkSaveSegment_updatedMapIDParallel benchmarks updating segment map IDs in parallel.
func (f Factory) BenchmarkSaveSegment_updatedMapIDParallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
				b.Error(err)
			}
		}
	})
}
