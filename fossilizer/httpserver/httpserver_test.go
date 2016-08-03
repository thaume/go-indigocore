package httpserver

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

	. "github.com/stratumn/go/fossilizer/adapter"
	"github.com/stratumn/go/fossilizer/adapter/adaptertest"
)

// Tests the root route if successful.
func TestRootOK(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockGetInfo.Fn = func() (interface{}, error) { return "test", nil }

	var dict map[string]interface{}
	res, err := getJSON(server.URL, &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("unexpected status code")
	}
	if dict["adapter"].(string) != "test" {
		t.Fatal("unexpected adapter dict")
	}
	if adapter.MockGetInfo.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetInfo()")
	}
}

// Tests the root route if an error occured in the adapter.
func TestRootErr(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockGetInfo.Fn = func() (interface{}, error) { return "test", errors.New("error") }

	var dict map[string]interface{}
	res, err := getJSON(server.URL, &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrInternalServer.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrInternalServer.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockGetInfo.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetInfo()")
	}
}

// Tests the fossilize route.
func TestFossilizeOK(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		t.Fatal(err)
	}

	handler := &ResultHandler{T: t, Listener: listener, Expected: "\"it is known\""}

	go func() {
		defer listener.Close()

		resultChan := adapter.MockAddResultChan.LastCalledWith

		adapter.MockFossilize.Fn = func(data []byte, meta []byte) error {
			resultChan <- &Result{
				Evidence: "it is known",
				Data:     data,
				Meta:     meta,
			}
			return nil
		}

		values := url.Values{}
		values.Set("data", "1234567890")
		values.Set("callbackUrl", "http://localhost:6666")
		res, err := http.PostForm(server.URL+"/fossils", values)

		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Fatal("unexpected status code")
		}

		time.Sleep(2 * time.Second)
		t.Fatal("callback URL not called")
	}()

	http.Serve(listener, handler)
}

// Tests the fossilize without data.
func TestFossilizeNoData(t *testing.T) {
	server, _ := createServer()
	defer server.Close()

	values := url.Values{}
	values.Set("callbackUrl", "http://localhost:6666")
	res, err := http.PostForm(server.URL+"/fossils", values)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatal("unexpected status code")
	}
}

// Tests the fossilize without a callback.
func TestFossilizeNoCallback(t *testing.T) {
	server, _ := createServer()
	defer server.Close()

	values := url.Values{}
	values.Set("data", "1234567890")
	res, err := http.PostForm(server.URL+"/fossils", values)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatal("unexpected status code")
	}
}

// Tests the fossilize without body.
func TestFossilizeNoBody(t *testing.T) {
	server, _ := createServer()
	defer server.Close()

	url := server.URL + "/fossils?callbackUrl=http%3A%2F%2Flocalhost%3A6666"
	res, err := http.Post(url, "application/octet-stream", nil)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatal("unexpected status code")
	}
}

// Tests the not found route.
func TestNotFound(t *testing.T) {
	server, _ := createServer()
	defer server.Close()

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/dsfsdf", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrNotFound.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrNotFound.Msg {
		t.Fatal("unexpected error message")
	}
}

func createServer() (*httptest.Server, *adaptertest.MockAdapter) {
	adapter := &adaptertest.MockAdapter{}
	server := httptest.NewServer(New(adapter, &Config{MinDataLen: 1}))

	return server, adapter
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
