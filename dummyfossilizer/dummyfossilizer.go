// Package dummyfossilizer implements a fossilizer that can be used for testing.
// It doesn't do much -- it just adds a timestamp, therefore is useless for other purposes.
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

// DummyFossilizer is the type that implements github.com/stratumn/go/fossilizer.Adapter.
type DummyFossilizer struct {
	version     string
	resultChans []chan *fossilizer.Result
}

// New creates an instance of a DummyFossilizer.
func New(version string) *DummyFossilizer {
	return &DummyFossilizer{version, nil}
}

// GetInfo implements github.com/stratumn/go/fossilizer.Adapter.GetInfo.
func (a *DummyFossilizer) GetInfo() (interface{}, error) {
	return map[string]interface{}{
		"name":        Name,
		"description": Description,
		"version":     a.version,
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

	for _, c1 := range a.resultChans {
		go func(c2 chan *fossilizer.Result) {
			c2 <- r
		}(c1)
	}

	return nil
}
