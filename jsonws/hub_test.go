// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package jsonws

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/stratumn/sdk/jsonws/jsonwstesting"
)

func TestHubBroadcast(t *testing.T) {
	h := NewHub()
	go h.Start()

	c1 := &jsonwstesting.MockConn{}
	c2 := &jsonwstesting.MockConn{}

	// Check errors don't stop the hub.
	c1.MockWriteJSON.Fn = func(interface{}) error {
		return errors.New("test")
	}

	h.Register(c1)
	h.Register(c2)

	m := map[string]string{"msg": "hello"}

	h.Broadcast(m, nil)
	h.Stop()

	if got, want := c1.MockWriteJSON.CalledCount, 1; got != want {
		t.Errorf(`c1.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}
	if got, want := c2.MockWriteJSON.CalledCount, 1; got != want {
		t.Errorf(`c2.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}

	if got, want := c1.MockWriteJSON.LastCalledWith, m; !reflect.DeepEqual(got, want) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(want, "", "  ")
		t.Errorf("c1.MockWriteJSON.LastCalledWith = %s\n want %s", gotJS, wantJS)
	}
	if got, want := c2.MockWriteJSON.LastCalledWith, m; !reflect.DeepEqual(got, want) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(want, "", "  ")
		t.Errorf("c2.MockWriteJSON.LastCalledWith = %s\n want %s", gotJS, wantJS)
	}
}

func TestHubUnregister(t *testing.T) {
	h := NewHub()
	go h.Start()

	c1 := &jsonwstesting.MockConn{}
	c2 := &jsonwstesting.MockConn{}

	h.Register(c1)
	h.Register(c2)
	h.Unregister(c1)

	m := map[string]string{"msg": "hello"}

	h.Broadcast(m, nil)
	h.Stop()

	if got, want := c1.MockWriteJSON.CalledCount, 0; got != want {
		t.Errorf(`c1.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}
	if got, want := c2.MockWriteJSON.CalledCount, 1; got != want {
		t.Errorf(`c2.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}
}

func TestHubTag(t *testing.T) {
	h := NewHub()
	go h.Start()

	c1 := &jsonwstesting.MockConn{}
	c2 := &jsonwstesting.MockConn{}

	h.Register(c1)
	h.Register(c2)

	h.Tag(c1, "test")

	m := map[string]string{"msg": "hello"}

	h.Broadcast(m, "test")

	// Check Unregister removes connection from tag. This call to Broadcast
	// should not reach any connection.
	h.Unregister(c1)
	h.Broadcast(m, "test")
	h.Stop()

	if got, want := c1.MockWriteJSON.CalledCount, 1; got != want {
		t.Errorf(`c1.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}
	if got, want := c2.MockWriteJSON.CalledCount, 0; got != want {
		t.Errorf(`c2.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}

	if got, want := c1.MockWriteJSON.LastCalledWith, m; !reflect.DeepEqual(got, want) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(want, "", "  ")
		t.Errorf("c1.MockWriteJSON.LastCalledWith = %s\n want %s", gotJS, wantJS)
	}
}

func TestHubUntag(t *testing.T) {
	h := NewHub()
	go h.Start()

	c1 := &jsonwstesting.MockConn{}
	c2 := &jsonwstesting.MockConn{}

	h.Register(c1)
	h.Register(c2)

	h.Tag(c1, "test")
	h.Tag(c2, "test")
	h.Untag(c2, "test")

	m := map[string]string{"msg": "hello"}

	h.Broadcast(m, "test")
	h.Stop()

	if got, want := c1.MockWriteJSON.CalledCount, 1; got != want {
		t.Errorf(`c1.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}
	if got, want := c2.MockWriteJSON.CalledCount, 0; got != want {
		t.Errorf(`c2.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}

	if got, want := c1.MockWriteJSON.LastCalledWith, m; !reflect.DeepEqual(got, want) {
		gotJS, _ := json.MarshalIndent(got, "", "  ")
		wantJS, _ := json.MarshalIndent(want, "", "  ")
		t.Errorf("c1.MockWriteJSON.LastCalledWith = %s\n want %s", gotJS, wantJS)
	}
}
