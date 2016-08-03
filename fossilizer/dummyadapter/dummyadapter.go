package dummyadapter

import (
	"time"

	. "github.com/stratumn/go/fossilizer/adapter"
)

const (
	NAME            = "dummy"
	DESCRIPTION     = "Stratumn Dummy Adapter"
	DEFAULT_VERSION = "0.1.0"
)

// The type of the dummy adapter.
type DummyAdapter struct {
	version     string
	resultChans []chan *Result
}

// Creates a new dummy adapter.
func New(version string) *DummyAdapter {
	if version == "" {
		version = DEFAULT_VERSION
	}

	return &DummyAdapter{version, nil}
}

// Implements github.com/stratumn/go/fossilizer/adapter.
func (a *DummyAdapter) GetInfo() (interface{}, error) {
	return map[string]interface{}{
		"name":        NAME,
		"description": DESCRIPTION,
		"version":     a.version,
	}, nil
}

// Implements github.com/stratumn/go/fossilizer/adapter.
func (a *DummyAdapter) AddResultChan(resultChan chan *Result) {
	a.resultChans = append(a.resultChans, resultChan)
}

// Implements github.com/stratumn/go/fossilizer/adapter.
func (a *DummyAdapter) Fossilize(data []byte, meta []byte) error {
	defer func() {
		result := &Result{
			Evidence: map[string]interface{}{
				"authority": "dummy",
				"timestamp": time.Now().UTC().Format("20060102150405"),
			},
			Data: data,
			Meta: meta,
		}

		for _, c := range a.resultChans {
			go func() {
				c <- result
			}()
		}
	}()

	return nil
}
