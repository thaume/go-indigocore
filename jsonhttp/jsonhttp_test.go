// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package jsonhttp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGet(t *testing.T) {
	s := New(&Config{})

	s.Get("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := getJSON(ts.URL+"/test", &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
	}
}

func TestPost(t *testing.T) {
	s := New(&Config{})

	s.Post("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := postJSON(ts.URL+"/test", &body, nil)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
	}
}

func TestPut(t *testing.T) {
	s := New(&Config{})

	s.Put("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := putJSON(ts.URL+"/test", &body, nil)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
	}
}

func TestDelete(t *testing.T) {
	s := New(&Config{})

	s.Delete("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := deleteJSON(ts.URL+"/test", &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
	}
}

func TestNotFound(t *testing.T) {
	s := New(&Config{})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]interface{}
	res, err := getJSON(ts.URL+"/test", &body)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != NewErrNotFound("").Status() {
		t.Fatal("unexpected HTTP status")
	}

	if body["error"].(string) != NewErrNotFound("").Error() {
		t.Fatal("unexpected error")
	}

	if int(body["status"].(float64)) != NewErrNotFound("").Status() {
		t.Fatal("unexpected error HTTP status")
	}
}

func TestError(t *testing.T) {
	s := New(&Config{})

	s.Get("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return nil, NewErrBadRequest("no")
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]interface{}
	res, err := getJSON(ts.URL+"/test", &body)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != NewErrBadRequest("").Status() {
		t.Fatal("unexpected HTTP status")
	}

	if body["error"].(string) != "no" {
		t.Fatal("unexpected error")
	}

	if int(body["status"].(float64)) != NewErrBadRequest("").Status() {
		t.Fatal("unexpected error HTTP status")
	}
}

func TestPatch(t *testing.T) {
	s := New(&Config{})

	s.Patch("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := patchJSON(ts.URL+"/test", &body, nil)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
	}
}

func TestOptions(t *testing.T) {
	s := New(&Config{})

	s.Options("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := optionsJSON(ts.URL+"/test", &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
	}
}

func getJSON(url string, target interface{}) (*http.Response, error) {
	return requestJSON(http.MethodGet, url, target, nil)
}

func postJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return requestJSON(http.MethodPost, url, target, payload)
}

func putJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return requestJSON(http.MethodPut, url, target, payload)
}

func deleteJSON(url string, target interface{}) (*http.Response, error) {
	return requestJSON(http.MethodDelete, url, target, nil)
}

func patchJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return requestJSON(http.MethodPatch, url, target, payload)
}

func optionsJSON(url string, target interface{}) (*http.Response, error) {
	return requestJSON(http.MethodOptions, url, target, nil)
}

func requestJSON(method, url string, target, payload interface{}) (*http.Response, error) {
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
