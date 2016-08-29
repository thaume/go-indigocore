// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package jsonhttp

import (
	"net/http"
	"testing"
)

func TestNewErrInternalServer(t *testing.T) {
	testErrStatus(t, NewErrInternalServer(""), http.StatusInternalServerError)
	testErrError(t, NewErrInternalServer(""), "internal server error")
	testErrError(t, NewErrInternalServer("test"), "test")
}

func TestNewErrBadRequest(t *testing.T) {
	testErrStatus(t, NewErrBadRequest(""), http.StatusBadRequest)
	testErrError(t, NewErrBadRequest(""), "bad request")
	testErrError(t, NewErrBadRequest("test"), "test")
}

func TestNewErrUnauthorized(t *testing.T) {
	testErrStatus(t, NewErrUnauthorized(""), http.StatusUnauthorized)
	testErrError(t, NewErrUnauthorized(""), "unauthorized")
	testErrError(t, NewErrUnauthorized("test"), "test")
}

func TestNewErrNotFound(t *testing.T) {
	testErrStatus(t, NewErrNotFound(""), http.StatusNotFound)
	testErrError(t, NewErrNotFound(""), "not found")
	testErrError(t, NewErrNotFound("test"), "test")
}
