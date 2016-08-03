package adaptertest

import (
	. "github.com/stratumn/go/fossilizer/adapter"
)

// A type to mock an adapter.
type MockAdapter struct {
	MockGetInfo       MockGetInfo
	MockAddResultChan MockAddResultChan
	MockFossilize     MockFossilize
}

type MockGetInfo struct {
	CalledCount int
	Fn          func() (interface{}, error)
}

type MockAddResultChan struct {
	CalledCount    int
	CalledWith     []chan *Result
	LastCalledWith chan *Result
	Fn             func(chan *Result)
}

type MockFossilize struct {
	CalledCount        int
	CalledWithData     [][]byte
	CalledWithMeta     [][]byte
	LastCalledWithData []byte
	LastCalledWithMeta []byte
	Fn                 func([]byte, []byte) error
}

// Implements github.com/stratumn/go/fossilizer/adapter.
func (a *MockAdapter) GetInfo() (interface{}, error) {
	a.MockGetInfo.CalledCount++

	if a.MockGetInfo.Fn != nil {
		return a.MockGetInfo.Fn()
	}

	return nil, nil
}

// Implements github.com/stratumn/go/fossilizer/adapter.
func (a *MockAdapter) AddResultChan(resultChan chan *Result) {
	a.MockAddResultChan.CalledCount++
	a.MockAddResultChan.CalledWith = append(a.MockAddResultChan.CalledWith, resultChan)
	a.MockAddResultChan.LastCalledWith = resultChan

	if a.MockAddResultChan.Fn != nil {
		a.MockAddResultChan.Fn(resultChan)
	}
}

// Implements github.com/stratumn/go/fossilizer/adapter.
func (a *MockAdapter) Fossilize(data []byte, meta []byte) error {
	a.MockFossilize.CalledCount++
	a.MockFossilize.CalledWithData = append(a.MockFossilize.CalledWithData, data)
	a.MockFossilize.LastCalledWithData = data
	a.MockFossilize.CalledWithMeta = append(a.MockFossilize.CalledWithMeta, meta)
	a.MockFossilize.LastCalledWithMeta = meta

	if a.MockFossilize.Fn != nil {
		return a.MockFossilize.Fn(data, meta)
	}

	return nil
}
