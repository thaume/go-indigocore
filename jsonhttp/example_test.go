// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

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

func ExampleServer() {
	// Create the server.
	s := jsonhttp.New(&jsonhttp.Config{Port: ":3333"})

	// Add a route with a named param.
	s.Get("/items/:id", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *jsonhttp.Config) (interface{}, error) {
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
