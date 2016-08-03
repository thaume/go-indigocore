package httpserver

import (
	"encoding/json"
)

// An error with an HTTP status.
type ErrHTTP struct {
	Msg    string `json:"error"`
	Status int    `json:"status"`
}

// Makes ErrHTTP comply to the error interface.
func (e *ErrHTTP) Error() string {
	return e.Msg
}

// Helper to convert an error to JSON.
func (e *ErrHTTP) JSONEncode() []byte {
	js, err := json.Marshal(e)

	if err != nil {
		msg := `{"error:": "an internal server error occured", "status": 500}`
		return []byte(msg)
	}

	return js
}
