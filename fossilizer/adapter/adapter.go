package adapter

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
