// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizerhttp

import (
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/stratumn/go/fossilizer/fossilizertesting"
	"github.com/stratumn/go/jsonhttp"
)

func createServer() (*jsonhttp.Server, *fossilizertesting.MockAdapter) {
	a := &fossilizertesting.MockAdapter{}
	s := New(a, &Config{MinDataLen: 2, MaxDataLen: 16})

	return s, a
}

type resultHandler struct {
	t        *testing.T
	listener net.Listener
	want     string
}

func (h *resultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.listener.Close()

	w.Write([]byte("thanks"))
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.t.Fatalf("ioutil.ReadAll(): err: %s", err)
		return
	}

	if got, want := string(body), h.want; got != want {
		h.t.Errorf("h.ServerHTTP(): body = %q want %q", got, want)
	}
}
