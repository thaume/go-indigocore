// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package fossilizertesting_test

import (
	"fmt"
	"log"

	"github.com/stratumn/go/fossilizer/fossilizertesting"
)

// This example shows how to use a mock adapter.
func ExampleMockAdapter() {
	// Create a mock.
	m := fossilizertesting.MockAdapter{}

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
