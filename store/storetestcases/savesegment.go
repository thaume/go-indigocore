// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"io/ioutil"
	"log"
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

// TestSaveSegmentUpdatedState tests what happens when you update the state of a segment.
func (f Factory) TestSaveSegmentUpdatedState(t *testing.T) {
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

// TestSaveSegmentUpdatedMapID tests what happens when you update the map ID of a segment.
func (f Factory) TestSaveSegmentUpdatedMapID(t *testing.T) {
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

// TestSaveSegmentBranch tests what happens when you save a segment with a previous link hash.
func (f Factory) TestSaveSegmentBranch(t *testing.T) {
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
	log.SetOutput(ioutil.Discard)

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Error(err)
		}
	}
}

// BenchmarkSaveSegmentParallel benchmarks saving new segments in parallel.
func (f Factory) BenchmarkSaveSegmentParallel(b *testing.B) {
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
	log.SetOutput(ioutil.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1
			if err := a.SaveSegment(slice[i]); err != nil {
				b.Error(err)
			}
		}
	})
}

// BenchmarkSaveSegmentUpdatedState benchmarks updating segments states.
func (f Factory) BenchmarkSaveSegmentUpdatedState(b *testing.B) {
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
	log.SetOutput(ioutil.Discard)

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Error(err)
		}
	}
}

// BenchmarkSaveSegmentUpdatedStateParallel benchmarks updating segments states in parallel.
func (f Factory) BenchmarkSaveSegmentUpdatedStateParallel(b *testing.B) {
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
	log.SetOutput(ioutil.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1
			if err := a.SaveSegment(slice[i]); err != nil {
				b.Error(err)
			}
		}
	})
}

// BenchmarkSaveSegmentUpdatedMapID benchmarks updating segment map IDs.
func (f Factory) BenchmarkSaveSegmentUpdatedMapID(b *testing.B) {
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
	log.SetOutput(ioutil.Discard)

	for i := 0; i < b.N; i++ {
		if err := a.SaveSegment(slice[i]); err != nil {
			b.Error(err)
		}
	}
}

// BenchmarkSaveSegmentUpdatedMapIDParallel benchmarks updating segment map IDs in parallel.
func (f Factory) BenchmarkSaveSegmentUpdatedMapIDParallel(b *testing.B) {
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
	log.SetOutput(ioutil.Discard)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := atomic.AddUint64(&counter, 1) - 1
			if err := a.SaveSegment(slice[i]); err != nil {
				b.Error(err)
			}
		}
	})
}
