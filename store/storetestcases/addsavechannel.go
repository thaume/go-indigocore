// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package storetestcases

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/cs/cstesting"
)

// TestAddDidSaveChannel tests that AddDidSaveChannel functions properly.
func (f Factory) TestAddDidSaveChannel(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatalf("f.New(): err: %s", err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	c := make(chan *cs.Segment, 1)
	a.AddDidSaveChannel(c)

	s := cstesting.RandomSegment()
	if err := a.SaveSegment(s); err != nil {
		t.Fatalf("a.SaveSegment(): err: %s", err)
	}

	if got, want := <-c, s; !reflect.DeepEqual(want, got) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(got, "", "  ")
		t.Errorf("<- c = %s\n want%s", gotJS, wantJS)
	}
}
