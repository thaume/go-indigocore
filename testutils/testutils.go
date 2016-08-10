// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package testutils contains helpers for tests.
package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetJSON(url string, target interface{}) (*http.Response, error) {
	return RequestJSON(http.MethodGet, url, target, nil)
}

func PostJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return RequestJSON(http.MethodPost, url, target, payload)
}

func PutJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return RequestJSON(http.MethodPut, url, target, payload)
}

func DeleteJSON(url string, target interface{}) (*http.Response, error) {
	return RequestJSON(http.MethodDelete, url, target, nil)
}

func PatchJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return RequestJSON(http.MethodPatch, url, target, payload)
}

func OptionsJSON(url string, target interface{}) (*http.Response, error) {
	return RequestJSON(http.MethodOptions, url, target, nil)
}

func RequestJSON(method, url string, target, payload interface{}) (*http.Response, error) {
	var req *http.Request
	var err error
	var body []byte

	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&target); err != nil {
		return nil, err
	}

	return res, nil
}
