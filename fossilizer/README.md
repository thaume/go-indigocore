# Stratumn fossilizer

A Golang package to create Stratumn fossilizers.

[![build status](https://travis-ci.org/stratumn/fossilizer.svg)](https://travis-ci.org/stratumn/fossilizer.svg)

## Adapters

An adapter must implement this interface:

```go
type Adapter interface {
	// Returns arbitrary information about the adapter.
	GetInfo() (interface{}, error)

	// Adds a channel that receives results whenever data is fossilized.
	AddResultChan(resultChan chan *Result)

	// Requests data to be fossilized.
	// Meta is arbitrary data that will be sent to the result channels.
	Fossilize(data []byte, meta []byte) error
}

type Result struct {
	Evidence interface{}
	Data     []byte
	Meta     []byte
}
```

You can then use `github.com/stratumn/go/fossilizer/httpserver` to create an HTTP server for that adapter.

See `github.com/stratumn/go/fossilizer/dummyadapter` for an example.
