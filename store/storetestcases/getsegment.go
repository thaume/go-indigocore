// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/stratumn/go/cs/cstesting"
)

// TestGetSegmentFound tests what happens when you get an existing segment.
func (f Factory) TestGetSegmentFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s1 := cstesting.RandomSegment()
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

// TestGetSegmentUpdatedState tests what happens when you get a segment whose state was updated.
func (f Factory) TestGetSegmentUpdatedState(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s1 := cstesting.RandomSegment()
	linkHash := s1.Meta["linkHash"].(string)

	a.SaveSegment(s1)
	s1 = cstesting.ChangeSegmentState(s1)
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

// TestGetSegmentUpdatedMapID tests what happens when you get a segment whose map ID was updated.
func (f Factory) TestGetSegmentUpdatedMapID(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s1 := cstesting.RandomSegment()
	linkHash := s1.Meta["linkHash"].(string)

	a.SaveSegment(s1)
	s1 = cstesting.ChangeSegmentMapID(s1)
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

// TestGetSegmentNotFound tests what happens when you get a nonexistent segment.
func (f Factory) TestGetSegmentNotFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s, err := a.GetSegment(cstesting.RandomString(32))

	if err != nil {
		t.Fatal(err)
	}

	if s != nil {
		t.Fatal("expected segment to be nil")
	}
}

// BenchmarkGetSegmentFound benchmarks getting existing segments.
func (f Factory) BenchmarkGetSegmentFound(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		s := cstesting.RandomSegment()
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

// BenchmarkGetSegmentFoundParallel benchmarks getting existing segments in parallel.
func (f Factory) BenchmarkGetSegmentFoundParallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	linkHashes := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		s := cstesting.RandomSegment()
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
