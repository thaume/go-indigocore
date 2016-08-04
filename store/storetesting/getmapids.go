// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetesting

import (
	"fmt"
	"testing"

	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/store"
)

// TestGetMapIDsAll tests what happens when you get all the map IDs.
func TestGetMapIDsAll(t *testing.T, a store.Adapter) {
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

	if len(slice) != 10 {
		t.Fatal("expected map length to be 10")
	}

	for i := 0; i < 10; i++ {
		if !cstesting.ContainsString(slice, fmt.Sprintf("map%d", i)) {
			t.Fatal("missing map ID")
		}
	}
}

// TestGetMapIDsPagination tests what happens when you get map IDs with pagination.
func TestGetMapIDsPagination(t *testing.T, a store.Adapter) {
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

	if len(slice) != 5 {
		t.Fatal("expected map length to be 5")
	}
}

// TestGetMapIDsEmpty tests what happens when you should get no map IDs.
func TestGetMapIDsEmpty(t *testing.T, a store.Adapter) {
	slice, err := a.GetMapIDs(&store.Pagination{Offset: 100000, Limit: 5})

	if err != nil {
		t.Fatal(err)
	}

	if len(slice) != 0 {
		t.Fatal("expected map length to be 0")
	}
}
