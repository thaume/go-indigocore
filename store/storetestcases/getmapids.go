// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"fmt"
	"testing"

	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/store"
	"github.com/stratumn/go/testutil"
)

// TestGetMapIDs tests what happens when you get all the map IDs.
func (f Factory) TestGetMapIDs(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			s := cstesting.RandomSegment()
			s.Link.Meta["mapId"] = fmt.Sprintf("map%d", i)
			a.SaveSegment(s)
		}
	}

	slice, err := a.GetMapIDs(&store.Pagination{})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 10; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}

	for i := 0; i < 10; i++ {
		mapID := fmt.Sprintf("map%d", i)
		if !testutil.ContainsString(slice, mapID) {
			t.Errorf("slice does not contain %q", mapID)
		}
	}
}

// TestGetMapIDsPagination tests what happens when you get map IDs with pagination.
func (f Factory) TestGetMapIDsPagination(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			segment := cstesting.RandomSegment()
			segment.Link.Meta["mapId"] = fmt.Sprintf("map%d", i)
			a.SaveSegment(segment)
		}
	}

	slice, err := a.GetMapIDs(&store.Pagination{Offset: 3, Limit: 5})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 5; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestGetMapIDsEmpty tests what happens when you should get no map IDs.
func (f Factory) TestGetMapIDsEmpty(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	slice, err := a.GetMapIDs(&store.Pagination{Offset: 100000, Limit: 5})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(slice), 0; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}
