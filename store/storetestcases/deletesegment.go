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

// TestDeleteSegmentFound tests what happens when you delete an existing segments.
func (f Factory) TestDeleteSegmentFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s1 := cstesting.RandomSegment()
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

// TestDeleteSegmentNotFound tests what happens when you delete a nonexistent segment.
func (f Factory) TestDeleteSegmentNotFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("expected adapter not to be nil")
	}
	defer f.free(a)

	s, err := a.DeleteSegment(cstesting.RandomString(32))

	if err != nil {
		t.Fatal(err)
	}

	if s != nil {
		t.Fatal("expected segment to be nil")
	}
}

// BenchmarkDeleteSegmentFound benchmarks deleting existing segments.
func (f Factory) BenchmarkDeleteSegmentFound(b *testing.B) {
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
		if s, err := a.DeleteSegment(linkHashes[i]); err != nil {
			b.Fatal(err)
		} else if s == nil {
			b.Fatal("expected segment")
		}
	}
}

// BenchmarkDeleteSegmentFoundParallel benchmarks deleting existing segments in parallel.
func (f Factory) BenchmarkDeleteSegmentFoundParallel(b *testing.B) {
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

			if s, err := a.DeleteSegment(linkHashes[i]); err != nil {
				b.Fatal(err)
			} else if s == nil {
				b.Fatal("expected segment")
			}
		}
	})
}
