// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package dummyfossilizer

import (
	"testing"

	"github.com/stratumn/go/fossilizer"
)

func TestFossilize(t *testing.T) {
	a := New("")
	rc := make(chan *fossilizer.Result)
	a.AddResultChan(rc)

	data := []byte("data")
	meta := []byte("meta")

	go func() {
		if err := a.Fossilize(data, meta); err != nil {
			t.Fatal(err)
		}
	}()

	r := <-rc

	if string(r.Data) != string(data) {
		t.Fatal("Unexpected result data")
	}

	if string(r.Meta) != string(meta) {
		t.Fatal("Unexpected result meta")
	}

	if r.Evidence.(map[string]interface{})["authority"].(string) != "dummy" {
		t.Fatal("Unexpected result evidence")
	}
}
