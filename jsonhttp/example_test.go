// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package jsonhttp_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"

	"github.com/stratumn/go/jsonhttp"
)

// This example shows how to create a server and add a route with a named param.
// It also tests the route using net/http/httptest.
func ExampleServer() {
	// Create the server.
	s := jsonhttp.New(&jsonhttp.Config{Address: ":3333"})

	// Add a route with a named param.
	s.Get("/items/:id", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params) (interface{}, error) {
		// Return a map containing the ID.
		result := map[string]string{
			"id": p.ByName("id"),
		}

		return result, nil
	})

	// Create a test server.
	ts := httptest.NewServer(s)
	defer ts.Close()

	// Test our route.
	res, err := http.Get(ts.URL + "/items/one")
	if err != nil {
		log.Fatal(err)
	}

	item, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", item)
	// Output: {"id":"one"}
}
