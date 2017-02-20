// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package jsonws

import (
	"errors"
	"testing"
	"time"

	"github.com/stratumn/sdk/jsonws/jsonwstesting"
)

func TestBufferedConnWriteJSON(t *testing.T) {
	c := &jsonwstesting.MockConn{}
	bc := NewBufferedConn(c, &BufferedConnConfig{
		Size:         2,
		WriteTimeout: time.Second,
		PongTimeout:  2 * time.Second,
		PingInterval: time.Second,
		MaxMsgSize:   1024,
	})
	defer bc.Close()

	go func() {
		if err := bc.Start(); err != nil {
			t.Errorf(`bc.Start(): err: %s`, err)
		}
	}()

	m := map[string]string{"msg": "hello"}

	bc.WriteJSON(m)
	bc.WriteJSON(m)
	bc.WriteJSON(m)

	timer := time.After(time.Second)

TestBufferedConnWriteJSONLoop:
	for {
		select {
		case <-timer:
			break TestBufferedConnWriteJSONLoop
		default:
			if c.MockWriteJSON.CalledCount > 2 {
				break TestBufferedConnWriteJSONLoop
			}
		}
	}

	if got, want := c.MockWriteJSON.CalledCount, 3; got != want {
		t.Errorf(`c.MockWriteJSON.CalledCount = %d want %d`, got, want)
	}
}

func TestBufferedConnWriteJSON_Error(t *testing.T) {
	c := &jsonwstesting.MockConn{}
	bc := NewBufferedConn(c, &BufferedConnConfig{
		Size:         2,
		WriteTimeout: time.Second,
		PongTimeout:  2 * time.Second,
		PingInterval: time.Second,
		MaxMsgSize:   1024,
	})

	c.MockWriteJSON.Fn = func(interface{}) error {
		return errors.New("test")
	}

	go func() {
		if err := bc.Start(); err == nil {
			t.Error(`bc.Start(): err = nil want error`)
		}
	}()

	m := map[string]string{"msg": "hello"}

	bc.WriteJSON(m)
}

func TestBufferedConnPing(t *testing.T) {
	c := &jsonwstesting.MockConn{}
	bc := NewBufferedConn(c, &BufferedConnConfig{
		Size:         2,
		WriteTimeout: time.Second,
		PongTimeout:  2 * time.Second,
		PingInterval: 100 * time.Millisecond,
		MaxMsgSize:   1024,
	})
	defer bc.Close()

	go func() {
		if err := bc.Start(); err != nil {
			t.Errorf(`bc.Start(): err: %s`, err)
		}
	}()

	timer := time.After(time.Second)

TestBufferedConnPingLoop:
	for {
		select {
		case <-timer:
			break TestBufferedConnPingLoop
		default:
			if c.MockPing.CalledCount > 0 {
				break TestBufferedConnPingLoop
			}
		}
	}

	if got, want := c.MockPing.CalledCount, 1; got != want {
		t.Errorf(`c.MockPing.CalledCount = %d want %d`, got, want)
	}
}

func TestBufferedConnPong(t *testing.T) {
	c := &jsonwstesting.MockConn{}
	bc := NewBufferedConn(c, &BufferedConnConfig{
		Size:         2,
		WriteTimeout: time.Second,
		PongTimeout:  2 * time.Second,
		PingInterval: time.Second,
		MaxMsgSize:   1024,
	})
	defer bc.Close()

	go func() {
		if err := bc.Start(); err != nil {
			t.Errorf(`bc.Start(): err: %s`, err)
		}
	}()

	timer := time.After(time.Second)

TestBufferedConnPongLoop1:
	for {
		select {
		case <-timer:
			t.Fatalf(`c.MockSetPongHandler.CalledCount = %d want 1`, c.MockSetPongHandler.CalledCount)
		default:
			if c.MockSetPongHandler.LastCalledWith != nil {
				break TestBufferedConnPongLoop1
			}
		}
	}

	c.MockSetPongHandler.LastCalledWith("")

	timer = time.After(time.Second)

TestBufferedConnPongLoop2:
	for {
		select {
		case <-timer:
			break TestBufferedConnPongLoop2
		default:
			if c.MockSetReadDeadline.CalledCount > 1 {
				break TestBufferedConnPongLoop2
			}
		}
	}

	if got, want := c.MockSetReadDeadline.CalledCount, 2; got != want {
		t.Errorf(`c.MockSetReadDeadline.CalledCount = %d want %d`, got, want)
	}
}
