// Copyright 2016 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storetestcases

import (
	"fmt"
	"testing"

	"github.com/stratumn/go/cs/cstesting"
	"github.com/stratumn/go/store"
	"github.com/stratumn/go/testutil"
)

// TestGetMapIDs tests what happens when you get map IDs with default pagination.
func (f Factory) TestGetMapIDs(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatalf("f.New(): err: %s", err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	for i := 0; i < store.DefaultLimit; i++ {
		for j := 0; j < store.DefaultLimit; j++ {
			s := cstesting.RandomSegment()
			s.Link.Meta["mapId"] = fmt.Sprintf("map%d", i)
			a.SaveSegment(s)
		}
	}

	slice, err := a.GetMapIDs(&store.Pagination{Limit: store.DefaultLimit * store.DefaultLimit})
	if err != nil {
		t.Fatalf("a.GetMapIDs(): err: %s", err)
	}

	if got, want := len(slice), store.DefaultLimit; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}

	for i := 0; i < store.DefaultLimit; i++ {
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
		t.Fatalf("f.New(): err: %s", err)
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
		t.Fatalf("a.GetMapIDs(): err: %s", err)
	}

	if got, want := len(slice), 5; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}

// TestGetMapIDsEmpty tests what happens when you should get no map IDs.
func (f Factory) TestGetMapIDsEmpty(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatalf("f.New(): err: %s", err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	slice, err := a.GetMapIDs(&store.Pagination{Offset: 100000, Limit: 5})
	if err != nil {
		t.Fatalf("a.GetMapIDs(): err: %s", err)
	}

	if got, want := len(slice), 0; want != got {
		t.Errorf("len(slice) = %d want %d", got, want)
	}
}
