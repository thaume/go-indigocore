// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package storetesting_test

import (
	"fmt"
	"github.com/stratumn/go/store/storetesting"
	"log"
)

// This example shows how to use a mock adapter.
func ExampleMockAdapter() {
	// Create a mock.
	m := storetesting.MockAdapter{}

	// Define a GetInfo function for our mock.
	m.MockGetInfo.Fn = func() (interface{}, error) {
		return map[string]string{
			"name": "test",
		}, nil
	}

	// Execute GetInfo on the mock.
	i, err := m.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	name := i.(map[string]string)["name"]

	// This is the number of times GetInfo was called.
	calledCount := m.MockGetInfo.CalledCount

	fmt.Printf("%s %d", name, calledCount)
	// Output: test 1
}
