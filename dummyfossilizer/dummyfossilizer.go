// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package dummyfossilizer implements a fossilizer that can be used for testing.
//
// It doesn't do much -- it just adds a timestamp.
package dummyfossilizer

import (
	"time"

	"github.com/stratumn/go/fossilizer"
)

const (
	// Name is the name set in the fossilizer's information.
	Name = "dummy"

	// Description is the description set in the fossilizer's information.
	Description = "Stratumn Dummy Fossilizer"
)

// Config contains configuration options for the store.
type Config struct {
	// A version string that will set in the store's information.
	Version string

	// A git commit hash that will set in the store's information.
	Commit string
}

// DummyFossilizer is the type that implements github.com/stratumn/go/fossilizer.Adapter.
type DummyFossilizer struct {
	config      *Config
	resultChans []chan *fossilizer.Result
}

// New creates an instance of a DummyFossilizer.
func New(config *Config) *DummyFossilizer {
	return &DummyFossilizer{config, nil}
}

// GetInfo implements github.com/stratumn/go/fossilizer.Adapter.GetInfo.
func (a *DummyFossilizer) GetInfo() (interface{}, error) {
	return map[string]interface{}{
		"name":        Name,
		"description": Description,
		"version":     a.config.Version,
		"commit":      a.config.Commit,
	}, nil
}

// AddResultChan implements github.com/stratumn/go/fossilizer.Adapter.AddResultChan.
func (a *DummyFossilizer) AddResultChan(resultChan chan *fossilizer.Result) {
	a.resultChans = append(a.resultChans, resultChan)
}

// Fossilize implements github.com/stratumn/go/fossilizer.Adapter.Fossilize.
func (a *DummyFossilizer) Fossilize(data []byte, meta []byte) error {
	r := &fossilizer.Result{
		Evidence: map[string]interface{}{
			"authority": "dummy",
			"timestamp": time.Now().UTC().Format("20060102150405"),
		},
		Data: data,
		Meta: meta,
	}

	go func(chans []chan *fossilizer.Result) {
		for _, c := range chans {
			c <- r
		}
	}(a.resultChans)

	return nil
}
