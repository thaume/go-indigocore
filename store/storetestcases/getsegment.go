// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"encoding/json"
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/testutil"
)

// TestGetSegment tests what happens when you get an existing segment.
func (f Factory) TestGetSegment(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	s1 := cstesting.RandomSegment()
	linkHash := s1.Meta["linkHash"].(string)
	a.SaveSegment(s1)

	s2, err := a.GetSegment(linkHash)
	if err != nil {
		t.Fatal(err)
	}

	if got := s2; got == nil {
		t.Error("s2 = nil want *cs.Segment")
	}
	if got, want := s2, s1; !reflect.DeepEqual(want, got) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("s2 = %s\n want%s", gotJS, wantJS)
	}
}

// TestGetSegment_updatedState tests what happens when you get a segment whose state was updated.
func (f Factory) TestGetSegment_updatedState(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
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

	if got := s2; got == nil {
		t.Error("s2 = nil want *cs.Segment")
	}
	if got, want := s2, s1; !reflect.DeepEqual(want, got) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("s2 = %s\n want%s", gotJS, wantJS)
	}
}

// TestGetSegment_updatedMapID tests what happens when you get a segment whose map ID was updated.
func (f Factory) TestGetSegment_updatedMapID(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
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

	if got := s2; got == nil {
		t.Error("s2 = nil want *cs.Segment")
	}
	if got, want := s2, s1; !reflect.DeepEqual(want, got) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("s2 = %s\n want%s", gotJS, wantJS)
	}
}

// TestGetSegment_notFound tests what happens when you get a nonexistent segment.
func (f Factory) TestGetSegment_notFound(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	s, err := a.GetSegment(testutil.RandomString(32))
	if err != nil {
		t.Fatal(err)
	}

	if got := s; got != nil {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("s = %s\n want nil", gotJS)
	}
}

// BenchmarkGetSegment benchmarks getting existing segments.
func (f Factory) BenchmarkGetSegment(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
			b.Error("s = nil want *cs.Segment")
		}
	}
}

// BenchmarkGetSegment_parallel benchmarks getting existing segments in parallel.
func (f Factory) BenchmarkGetSegment_parallel(b *testing.B) {
	a, err := f.New()
	if err != nil {
		b.Fatal(err)
	}
	if a == nil {
		b.Fatal("a = nil want store.Adapter")
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
				b.Error(err)
			} else if s == nil {
				b.Error("s = nil want *cs.Segment")
			}
		}
	})
}
