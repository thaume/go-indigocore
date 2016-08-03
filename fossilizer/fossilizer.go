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

// Result is the type sent the the result channels.
type Result struct {
	// Evidence created by the fossilizer.
	Evidence interface{}
	// The data that was fossilized.
	Data []byte
	// The meta data that was given to Adapter.Fossilize.
	Meta []byte
}
