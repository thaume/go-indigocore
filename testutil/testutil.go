// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package testutil contains helpers for tests.
package testutil

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"

	"github.com/stratumn/go/types"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomHash creates a random hash.
func RandomHash() *types.Bytes32 {
	var hash types.Bytes32
	for i := range hash {
		hash[i] = byte(letters[rand.Intn(len(letters))])
	}
	return &hash
}

// RandomString generates a random string.
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ContainsString checks if an array contains a string.
func ContainsString(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// RequestJSON does a request expecting a JSON response.
func RequestJSON(h http.HandlerFunc, method, target string, payload, dst interface{}) (*httptest.ResponseRecorder, error) {
	var req *http.Request

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		req = httptest.NewRequest(method, target, bytes.NewReader(body))
	} else {
		req = httptest.NewRequest(method, target, nil)
	}

	w := httptest.NewRecorder()
	h(w, req)

	if dst != nil {
		if err := json.NewDecoder(w.Body).Decode(&dst); err != nil {
			return nil, err
		}
	}

	return w, nil
}
