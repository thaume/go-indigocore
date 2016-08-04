// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetesting

import (
	"reflect"
	"testing"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/cs/cstesting"
)

func TestMockAdapter(t *testing.T) {
	a := &MockAdapter{}
	s1 := cstesting.RandomSegment()
	a.MockGetSegment.Fn = func(linkHash string) (*cs.Segment, error) { return s1, nil }
	s2, err := a.GetSegment("abcdef")

	if s1 != s2 {
		t.Fatal("expected segments to be equal")
	}

	if err != nil {
		t.Fatal("unexpected error")
	}

	a.GetSegment("ghij")

	if a.MockGetSegment.CalledCount != 2 {
		t.Fatal("unexpected MockGetSegment.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockGetSegment.CalledWith, []string{"abcdef", "ghij"}) {
		t.Fatal("unexpected MockGetSegment.LastCalledWith value")
	}

	if a.MockGetSegment.LastCalledWith != "ghij" {
		t.Fatal("unexpected MockGetSegment.LastCalledWith value")
	}
}
