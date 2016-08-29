// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package dummyfossilizer

import (
	"testing"

	"github.com/stratumn/go/fossilizer"
)

func TestGetInfo(t *testing.T) {
	a := New(&Config{})
	got, err := a.GetInfo()
	if err != nil {
		t.Fatal(err)
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
			t.Error(err)
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
