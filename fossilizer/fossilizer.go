// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package fossilizer defines types to implement a fossilizer.
package fossilizer

// Adapter must be implemented by a fossilier.
type Adapter interface {
	// Returns arbitrary information about the adapter.
	GetInfo() (interface{}, error)

	// Adds a channel that receives results whenever data is fossilized.
	AddResultChan(resultChan chan *Result)

	// Requests data to be fossilized.
	// Meta is arbitrary data that will be sent to the result channels.
	Fossilize(data []byte, meta []byte) error
}

// Result is the type sent to the result channels.
type Result struct {
	// Evidence created by the fossilizer.
	Evidence interface{}

	// The data that was fossilized.
	Data []byte

	// The meta data that was given to Adapter.Fossilize.
	Meta []byte
}
