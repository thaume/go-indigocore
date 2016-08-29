// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetestcases

import (
	"testing"
)

// TestGetInfo tests what happens when you get information about the adapter.
func (f Factory) TestGetInfo(t *testing.T) {
	a, err := f.New()
	if err != nil {
		t.Fatalf("f.New(): err: %s", err)
	}
	if a == nil {
		t.Fatal("a = nil want store.Adapter")
	}
	defer f.free(a)

	info, err := a.GetInfo()
	if err != nil {
		t.Fatalf("a.GetInfo(): err: %s", err)
	}
	if info == nil {
		t.Fatal("info = nil want interface{}")
	}
}
