// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizertesting

import (
	"reflect"
	"testing"

	"github.com/stratumn/go/fossilizer"
)

func TestMockAdapter_GetInfo(t *testing.T) {
	a := &MockAdapter{}

	_, err := a.GetInfo()

	if err != nil {
		t.Fatal("unexpected error")
	}

	a.MockGetInfo.Fn = func() (interface{}, error) { return map[string]string{"name": "test"}, nil }
	info, err := a.GetInfo()

	if err != nil {
		t.Fatal("unexpected error")
	}

	if info.(map[string]string)["name"] != "test" {
		t.Fatal("unexpect info")
	}

	if a.MockGetInfo.CalledCount != 2 {
		t.Fatal("unexpected MockGetInfo.CalledCount value")
	}
}

func TestMockAdapter_AddResultChan(t *testing.T) {
	a := &MockAdapter{}

	c := make(chan *fossilizer.Result)
	a.AddResultChan(c)

	if a.MockAddResultChan.CalledCount != 1 {
		t.Fatal("unexpected MockAddResultChan.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockAddResultChan.CalledWith, []chan *fossilizer.Result{c}) {
		t.Fatal("unexpected MockAddResultChan.LastCalledWith value")
	}

	if a.MockAddResultChan.LastCalledWith != c {
		t.Fatal("unexpected MockAddResultChan.LastCalledWith value")
	}
}

func TestMockAdapter_Fossilize(t *testing.T) {
	a := &MockAdapter{}

	d := []byte("data")
	m := []byte("meta")

	err := a.Fossilize(d, m)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if a.MockFossilize.CalledCount != 1 {
		t.Fatal("unexpected MockFossilize.CalledCount value")
	}

	if !reflect.DeepEqual(a.MockFossilize.CalledWithData, [][]byte{d}) {
		t.Fatal("unexpected MockFossilize.CalledWithData value")
	}

	if string(a.MockFossilize.LastCalledWithData) != string(d) {
		t.Fatal("unexpected MockFossilize.LastCalledWithData value")
	}

	if !reflect.DeepEqual(a.MockFossilize.CalledWithMeta, [][]byte{m}) {
		t.Fatal("unexpected MockFossilize.CalledWithMeta value")
	}

	if string(a.MockFossilize.LastCalledWithMeta) != string(m) {
		t.Fatal("unexpected MockFossilize.LastCalledWithMeta value")
	}
}
