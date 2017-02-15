// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package fossilizerhttp_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/stratumn/go/dummyfossilizer"
	"github.com/stratumn/go/fossilizer/fossilizerhttp"
	"github.com/stratumn/go/jsonhttp"
)

// This example shows how to create a server from a dummyfossilizer.
// It also tests the root route of the server using net/http/httptest.
func Example() {
	// Create a dummy adapter.
	a := dummyfossilizer.New(&dummyfossilizer.Config{Version: "0.1.0", Commit: "abc"})
	config := &fossilizerhttp.Config{
		MaxDataLen: 64,
	}
	httpConfig := &jsonhttp.Config{
		Address: ":6000",
	}

	// Create a server.
	s := fossilizerhttp.New(a, config, httpConfig)

	// Create a test server.
	ts := httptest.NewServer(s)
	defer ts.Close()

	// Test the root route.
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	info, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", info)
	// Output: {"adapter":{"name":"dummy","description":"Stratumn Dummy Fossilizer","version":"0.1.0","commit":"abc"}}
}
