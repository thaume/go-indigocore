// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storehttp_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/stratumn/go/dummystore"
	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/storehttp"
)

// This example shows how to create a server from a dummystore.
// It also tests the root route of the server using net/http/httptest.
func Example() {
	// Create a dummy adapter.
	a := dummystore.New(&dummystore.Config{Version: "0.1.0", Commit: "abc"})
	c := &jsonhttp.Config{
		Port: "5555",
	}

	// Create a server.
	s := storehttp.New(a, c)

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
	// Output: {"adapter":{"name":"dummy","description":"Stratumn Dummy Store","version":"0.1.0","commit":"abc"}}
}
