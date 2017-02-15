// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package dummyfossilizer

import (
	"testing"

	"github.com/stratumn/go/fossilizer"
)

func TestGetInfo(t *testing.T) {
	a := New(&Config{})
	got, err := a.GetInfo()
	if err != nil {
		t.Fatalf("a.GetInfo(): err: %s", err)
	}
	if _, ok := got.(*Info); !ok {
		t.Errorf("a.GetInfo(): info = %#v want *Info", got)
	}
}

func TestFossilize(t *testing.T) {
	a := New(&Config{})
	rc := make(chan *fossilizer.Result)
	a.AddResultChan(rc)

	var (
		data = []byte("data")
		meta = []byte("meta")
	)

	go func() {
		if err := a.Fossilize(data, meta); err != nil {
			t.Errorf("a.Fossilize(): err: %s", err)
		}
	}()

	r := <-rc

	if got, want := string(r.Data), string(data); got != want {
		t.Errorf("<-rc: Data = %q want %q", got, want)
	}
	if got, want := string(r.Meta), string(meta); got != want {
		t.Errorf("<-rc: Meta = %q want %q", got, want)
	}
	if got, want := r.Evidence.(map[string]interface{})["authority"].(string), "dummy"; got != want {
		t.Errorf(`<-rc: Evidence["authority"] = %q want %q`, got, want)
	}
}
