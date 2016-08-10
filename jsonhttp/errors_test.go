// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package jsonhttp

import (
	"net/http"
	"testing"
)

func TestNewErrInternalServer(t *testing.T) {
	if NewErrInternalServer("").Status() != http.StatusInternalServerError {
		t.Fatal("unexpected error HTTP status")
	}
	if NewErrInternalServer("").Error() != "internal server error" {
		t.Fatal("unexpected error")
	}
	if NewErrInternalServer("test").Error() != "test" {
		t.Fatal("unexpected error")
	}
}

func TestNewErrBadRequest(t *testing.T) {
	if NewErrBadRequest("").Status() != http.StatusBadRequest {
		t.Fatal("unexpected error HTTP status")
	}
	if NewErrBadRequest("").Error() != "bad request" {
		t.Fatal("unexpected error")
	}
	if NewErrBadRequest("test").Error() != "test" {
		t.Fatal("unexpected error")
	}
}

func TestNewErrUnauthorized(t *testing.T) {
	if NewErrUnauthorized("").Status() != http.StatusUnauthorized {
		t.Fatal("unexpected error HTTP status")
	}
	if NewErrUnauthorized("").Error() != "unauthorized" {
		t.Fatal("unexpected error")
	}
	if NewErrUnauthorized("test").Error() != "test" {
		t.Fatal("unexpected error")
	}
}

func TestNewErrNotFound(t *testing.T) {
	if NewErrNotFound("").Status() != http.StatusNotFound {
		t.Fatal("unexpected error HTTP status")
	}
	if NewErrNotFound("").Error() != "not found" {
		t.Fatal("unexpected error")
	}
	if NewErrNotFound("test").Error() != "test" {
		t.Fatal("unexpected error")
	}
}
