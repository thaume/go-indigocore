// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
