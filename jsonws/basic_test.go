// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package jsonws

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stratumn/sdk/jsonws/jsonwstesting"
)

func testUpgradeHandle(w http.ResponseWriter, r *http.Request, h http.Header) (PingableConn, error) {
	conn := jsonwstesting.MockConn{}
	read := false
	conn.MockReadJSON.Fn = func(v interface{}) error {
		if read {
			return io.EOF
		}
		read = true
		m := v.(*interface{})
		*m = map[string]string{"msg": "test"}
		return nil
	}
	return &conn, nil
}

func testMsgAllocator(msg *interface{}) {
	*msg = map[string]string{}
}

func TestBasicAddConnChannel(t *testing.T) {
	ws := NewBasic(&BasicConfig{
		UpgradeHandle: testUpgradeHandle,
	}, &BufferedConnConfig{
		PingInterval: time.Second,
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ws", nil)
	c := make(chan *BufferedConn)
	ws.AddConnChannel(c)

	go ws.Handle(w, r)

	select {
	case got := <-c:
		if got == nil {
			t.Errorf("<-c = nil want not nil")
		}
	case <-time.After(time.Second):
		t.Errorf("no connection sent to channel")
	}
}

func TestBasicAddMsgChannel(t *testing.T) {
	ws := NewBasic(&BasicConfig{
		UpgradeHandle: testUpgradeHandle,
		MsgAllocator:  testMsgAllocator,
	}, &BufferedConnConfig{
		PingInterval: time.Second,
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ws", nil)
	c := make(chan BasicConnMsg)
	ws.AddMsgChannel(c)

	go ws.Start()
	go ws.Handle(w, r)
	defer ws.Stop()

	select {
	case msg := <-c:
		if got, want := msg.Msg.(map[string]string)["msg"], "test"; got != want {
			t.Errorf(`msg = %q want %q`, got, want)
		}
	case <-time.After(time.Second):
		t.Errorf("no message sent to channel")
	}
}
