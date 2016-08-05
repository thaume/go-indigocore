// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package jsonhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	// ErrInternalServer is an error for when an internal server occurs.
	ErrInternalServer = NewErrHTTP("internal server error", http.StatusInternalServerError)

	// ErrBadRequest is an error for when a bad request occurs.
	ErrBadRequest = NewErrHTTP("bad request", http.StatusBadRequest)

	// ErrUnauthorized is an error for when an unauthorized request occurs.
	ErrUnauthorized = NewErrHTTP("unauthorized", http.StatusUnauthorized)

	// ErrNotFound is an error for when something isn't found.
	ErrNotFound = NewErrHTTP("not found", http.StatusNotFound)
)

// ErrHTTP is an error with an HTTP status.
type ErrHTTP struct {
	msg    string
	status int
}

// NewErrHTTP creates a new error with a message and HTTP status.
func NewErrHTTP(msg string, status int) ErrHTTP {
	return ErrHTTP{msg, status}
}

// Status returns the HTTP status of the error.
func (e ErrHTTP) Status() int {
	return e.status
}

// Error implements error.Error.
func (e ErrHTTP) Error() string {
	return e.msg
}

var internalServerJSON = fmt.Sprintf(`{"error:":"internal server error","status":%d}`, http.StatusInternalServerError)

// JSONEncode marshals an error to JSON.
func (e ErrHTTP) JSONEncode() []byte {
	js, err := json.Marshal(map[string]interface{}{
		"error":  e.msg,
		"status": e.status,
	})

	if err != nil {
		msg := internalServerJSON
		return []byte(msg)
	}

	return js
}
