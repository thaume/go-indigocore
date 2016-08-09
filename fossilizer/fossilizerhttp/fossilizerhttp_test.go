// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package fossilizerhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stratumn/go/fossilizer"
	"github.com/stratumn/go/fossilizer/fossilizertesting"
	"github.com/stratumn/go/jsonhttp"
)

func TestRootOK(t *testing.T) {
	s, a := createServer()
	defer s.Close()

	a.MockGetInfo.Fn = func() (interface{}, error) { return "test", nil }

	var dict map[string]interface{}
	res, err := getJSON(s.URL, &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	if dict["adapter"].(string) != "test" {
		t.Fatal("unexpected adapter dict")
	}
	if a.MockGetInfo.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetInfo()")
	}
}

func TestRootErr(t *testing.T) {
	s, a := createServer()
	defer s.Close()

	a.MockGetInfo.Fn = func() (interface{}, error) { return "test", errors.New("error") }

	var dict map[string]interface{}
	res, err := getJSON(s.URL, &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != jsonhttp.NewErrInternalServer("").Status() {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != jsonhttp.NewErrInternalServer("").Error() {
		t.Fatal("unexpected error message")
	}
	if a.MockGetInfo.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetInfo()")
	}
}

func TestFossilizeOK(t *testing.T) {
	s, a := createServer()
	defer s.Close()

	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		t.Fatal(err)
	}

	h := &ResultHandler{T: t, Listener: l, Expected: "\"it is known\""}

	go func() {
		defer l.Close()

		rc := a.MockAddResultChan.LastCalledWith

		a.MockFossilize.Fn = func(data []byte, meta []byte) error {
			rc <- &fossilizer.Result{
				Evidence: "it is known",
				Data:     data,
				Meta:     meta,
			}
			return nil
		}

		v := url.Values{}
		v.Set("data", "1234567890")
		v.Set("callbackUrl", "http://localhost:6666")
		res, err := http.PostForm(s.URL+"/fossils", v)

		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatal("unexpected status code")
		}

		time.Sleep(2 * time.Second)
		t.Fatal("callback URL not called")
	}()

	http.Serve(l, h)
}

func TestFossilizeNoData(t *testing.T) {
	s, _ := createServer()
	defer s.Close()

	v := url.Values{}
	v.Set("callbackUrl", "http://localhost:6666")
	res, err := http.PostForm(s.URL+"/fossils", v)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("unexpected status code")
	}
}

func TestFossilizeNoCallback(t *testing.T) {
	s, _ := createServer()
	defer s.Close()

	v := url.Values{}
	v.Set("data", "1234567890")
	res, err := http.PostForm(s.URL+"/fossils", v)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("unexpected status code")
	}
}

func TestFossilizeNoBody(t *testing.T) {
	s, _ := createServer()
	defer s.Close()

	url := s.URL + "/fossils?callbackUrl=http%3A%2F%2Flocalhost%3A6666"
	res, err := http.Post(url, "application/octet-stream", nil)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Fatal("unexpected status code")
	}
}

func TestNotFound(t *testing.T) {
	s, _ := createServer()
	defer s.Close()

	var dict map[string]interface{}
	res, err := getJSON(s.URL+"/dsfsdf", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != jsonhttp.NewErrNotFound("").Status() {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != jsonhttp.NewErrNotFound("").Error() {
		t.Fatal("unexpected error message")
	}
}

func createServer() (*httptest.Server, *fossilizertesting.MockAdapter) {
	a := &fossilizertesting.MockAdapter{}
	s := httptest.NewServer(New(a, &Config{MinDataLen: 1}))

	return s, a
}

func getJSON(url string, target interface{}) (*http.Response, error) {
	return requestJSON(http.MethodGet, url, target, nil)
}

func postJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return requestJSON(http.MethodPost, url, target, payload)
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

type ResultHandler struct {
	T        *testing.T
	Listener net.Listener
	Expected string
}

func (h *ResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.Listener.Close()

	w.Write([]byte("thanks"))

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.T.Fatal(err)
	}

	if string(body) != h.Expected {
		h.T.Fatal("unexpected body")
	}
}
