// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package jsonhttp

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/stratumn/go/testutils"
)

func TestGet(t *testing.T) {
	s := New(&Config{})

	s.Get("/test", func(r http.ResponseWriter, _ *http.Request, p httprouter.Params, _ *Config) (interface{}, error) {
		return map[string]bool{"test": true}, nil
	})

	ts := httptest.NewServer(s)
	defer ts.Close()

	var body map[string]bool
	_, err := testutils.GetJSON(ts.URL+"/test", &body)
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
	_, err := testutils.PostJSON(ts.URL+"/test", &body, nil)
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
	_, err := testutils.PutJSON(ts.URL+"/test", &body, nil)
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
	_, err := testutils.DeleteJSON(ts.URL+"/test", &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, map[string]bool{"test": true}) {
		t.Fatal("unexpected body")
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
	_, err := testutils.PatchJSON(ts.URL+"/test", &body, nil)
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
	_, err := testutils.OptionsJSON(ts.URL+"/test", &body)
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
	res, err := testutils.GetJSON(ts.URL+"/test", &body)
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
	res, err := testutils.GetJSON(ts.URL+"/test", &body)
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
