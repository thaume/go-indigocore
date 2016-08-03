package httpserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/segment/segmenttest"
	. "github.com/stratumn/go/store/adapter"
	"github.com/stratumn/go/store/adapter/adaptertest"
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

// Tests the save segment route if the segment was successful.
func TestSaveSegmentOK(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockSaveSegment.Fn = func(*Segment) error { return nil }

	segment1 := RandomSegment()
	var segment2 Segment
	res, err := postJSON(server.URL+"/segments", &segment2, segment1)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockSaveSegment.LastCalledWith, segment1) {
		t.Fatal("unexpected argument passed to SaveSegment()")
	}
	if res.StatusCode != 200 {
		t.Fatal("unexpected status code")
	}
	if !reflect.DeepEqual(segment1, &segment2) {
		t.Fatal("expected segments to be equal")
	}
	if adapter.MockSaveSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to SaveSegment()")
	}
}

// Tests the save segment route if an error occured in the adapter.
func TestSaveSegmentErr(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockSaveSegment.Fn = func(*Segment) error { return errors.New("test") }

	var dict map[string]interface{}
	res, err := postJSON(server.URL+"/segments", &dict, RandomSegment())

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrInternalServer.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrInternalServer.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockSaveSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to SaveSegment()")
	}
}

// Tests the save segment route if a segment validation error occured.
func TestSaveSegmentInvalidSegment(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	segment := RandomSegment()
	segment.Meta["linkHash"] = true

	var dict map[string]interface{}
	res, err := postJSON(server.URL+"/segments", &dict, segment)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrBadRequest.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != "meta.linkHash should be a non empty string" {
		t.Fatal("unexpected error message")
	}
	if adapter.MockSaveSegment.CalledCount != 0 {
		t.Fatal("unexpected number of calls to SaveSegment()")
	}
}

// Tests the save segment route if a JSON error error occured.
func TestSaveSegmentInvalidJSON(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	var dict map[string]interface{}
	res, err := postJSON(server.URL+"/segments", &dict, "1234567890azertyui")

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrBadRequest.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrBadRequest.Msg {
		t.Log(dict["error"].(string))
		t.Log(ErrBadRequest.Msg)
		t.Fatal("unexpected error message")
	}
	if adapter.MockSaveSegment.CalledCount != 0 {
		t.Fatal("unexpected number of calls to SaveSegment()")
	}
}

// Tests the get segment route if the segment was found.
func TestGetSegmentFound(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	segment1 := RandomSegment()
	adapter.MockGetSegment.Fn = func(string) (*Segment, error) { return segment1, nil }

	var segment2 Segment
	res, err := getJSON(server.URL+"/segments/abcde", &segment2)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockGetSegment.LastCalledWith, "abcde") {
		t.Fatal("unexpected argument passed to GetSegment()")
	}
	if res.StatusCode != 200 {
		t.Fatal("unexpected status code")
	}
	if !reflect.DeepEqual(segment1, &segment2) {
		t.Fatal("expected segments to be equal")
	}
	if adapter.MockGetSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetSegment()")
	}
}

// Tests the get segment route if the segment was not found.
func TestGetSegmentNotFound(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/segments/abcde", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockGetSegment.LastCalledWith, "abcde") {
		t.Fatal("unexpected argument passed to GetSegment()")
	}
	if res.StatusCode != ErrNotFound.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrNotFound.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockGetSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetSegment()")
	}
}

// Tests the get segment route if an error occured in the adapter.
func TestGetSegmentErr(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockGetSegment.Fn = func(string) (*Segment, error) { return nil, errors.New("error") }

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/segments/abcde", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockGetSegment.LastCalledWith, "abcde") {
		t.Fatal("unexpected argument passed to GetSegment()")
	}
	if res.StatusCode != ErrInternalServer.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrInternalServer.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockGetSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to GetSegment()")
	}
}

// Tests the delete segment route if the segment was found.
func TestDeleteSegmentFound(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	segment1 := RandomSegment()
	adapter.MockDeleteSegment.Fn = func(string) (*Segment, error) { return segment1, nil }

	var segment2 Segment
	res, err := deleteJSON(server.URL+"/segments/abcde", &segment2)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockDeleteSegment.LastCalledWith, "abcde") {
		t.Fatal("unexpected argument passed to DeleteSegment()")
	}
	if res.StatusCode != 200 {
		t.Fatal("unexpected status code")
	}
	if !reflect.DeepEqual(segment1, &segment2) {
		t.Fatal("expected segments to be equal")
	}
	if adapter.MockDeleteSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to DeleteSegment()")
	}
}

// Tests the delete segment route if the segment was not found.
func TestDeleteSegmentNotFound(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	var dict map[string]interface{}
	res, err := deleteJSON(server.URL+"/segments/abcde", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockDeleteSegment.LastCalledWith, "abcde") {
		t.Fatal("unexpected argument passed to DeleteSegment()")
	}
	if res.StatusCode != ErrNotFound.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrNotFound.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockDeleteSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to DeleteSegment()")
	}
}

// Tests the delete segment route if an error occured in the adapter.
func TestDeleteSegmentErr(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockDeleteSegment.Fn = func(string) (*Segment, error) { return nil, errors.New("error") }

	var dict map[string]interface{}
	res, err := deleteJSON(server.URL+"/segments/abcde", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(adapter.MockDeleteSegment.LastCalledWith, "abcde") {
		t.Fatal("unexpected argument passed to DeleteSegment()")
	}
	if res.StatusCode != ErrInternalServer.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrInternalServer.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockDeleteSegment.CalledCount != 1 {
		t.Fatal("unexpected number of calls to DeleteSegment()")
	}
}

// Tests the get segment route if successful.
func TestFindSegmentsOK(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	var segments1 SegmentSlice
	for i := 0; i < 10; i++ {
		segments1 = append(segments1, RandomSegment())
	}
	adapter.MockFindSegments.Fn = func(*Filter) (SegmentSlice, error) { return segments1, nil }

	var segments2 SegmentSlice
	res, err := getJSON(server.URL+"/segments?offset=1&limit=2&mapId=123&prevLinkHash=abc&tags=one+two", &segments2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("unexpected status code")
	}
	if !reflect.DeepEqual(segments1, segments2) {
		t.Fatal("expected segment slices to be equal")
	}
	if adapter.MockFindSegments.CalledCount != 1 {
		t.Fatal("unexpected number of calls to FindSegments()")
	}

	filter := adapter.MockFindSegments.LastCalledWith
	if filter.Offset != 1 {
		t.Fatal("unexpected offset")
	}
	if filter.Limit != 2 {
		t.Fatal("unexpected limit")
	}
	if filter.MapID != "123" {
		t.Fatal("unexpected map ID")
	}
	if filter.PrevLinkHash != "abc" {
		t.Fatal("unexpected previous link hash")
	}
	if !reflect.DeepEqual(filter.Tags, []string{"one", "two"}) {
		t.Fatal("unexpected tags")
	}
}

// Tests the get segment route if an error occured in the adapter.
func TestFindSegmentsErr(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockFindSegments.Fn = func(*Filter) (SegmentSlice, error) { return nil, errors.New("test") }

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/segments?offset=1&limit=2&mapId=123&prevLinkHash=abc&tags=one,two", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrInternalServer.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrInternalServer.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockFindSegments.CalledCount != 1 {
		t.Fatal("unexpected number of calls to FindSegments()")
	}
}

// Tests the get segment route if an error occured in the query.
func TestFindSegmentsValidation(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/segments?offset=hello", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrOffset.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrOffset.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockFindSegments.CalledCount != 0 {
		t.Fatal("unexpected number of calls to FindSegments()")
	}
}

// Tests the get map IDs route if successful.
func TestGetMapIDsOK(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	slice1 := []string{"one", "two", "three"}
	adapter.MockGetMapIDs.Fn = func(*Pagination) ([]string, error) { return slice1, nil }

	var slice2 []string
	res, err := getJSON(server.URL+"/maps?offset=20&limit=10", &slice2)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatal("unexpected status code")
	}
	if len(slice1) != len(slice2) {
		t.Fatal("expected map ID slices to be have same length")
	}
	for i := 0; i < len(slice1); i++ {
		if !ContainsString(slice2, slice1[i]) {
			t.Fatal("expected map ID slices to have same elements")
		}
	}
	if adapter.MockGetMapIDs.CalledCount != 1 {
		t.Fatal("unexpected number of calls to FindSegments()")
	}

	pagination := adapter.MockGetMapIDs.LastCalledWith
	if pagination.Offset != 20 {
		t.Fatal("unexpected offset")
	}
	if pagination.Limit != 10 {
		t.Fatal("unexpected limit")
	}
}

// Tests the get map IDs route if an error occured in the adapter.
func TestGetMapIDsErr(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	adapter.MockGetMapIDs.Fn = func(*Pagination) ([]string, error) { return nil, errors.New("test") }

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/maps", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrInternalServer.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrInternalServer.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockGetMapIDs.CalledCount != 1 {
		t.Fatal("unexpected number of calls to FindSegments()")
	}
}

// Tests the get segment route if an error occured in the query.
func TestGetMapIDsValidation(t *testing.T) {
	server, adapter := createServer()
	defer server.Close()

	var dict map[string]interface{}
	res, err := getJSON(server.URL+"/maps?limit=-1", &dict)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != ErrLimit.Status {
		t.Fatal("unexpected status code")
	}
	if dict["error"].(string) != ErrLimit.Msg {
		t.Fatal("unexpected error message")
	}
	if adapter.MockGetMapIDs.CalledCount != 0 {
		t.Fatal("unexpected number of calls to FindSegments()")
	}
}

// Tests the not found route.
func TestRootNotFound(t *testing.T) {
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
	server := httptest.NewServer(New(adapter, &Config{}))

	return server, adapter
}

func getJSON(url string, target interface{}) (*http.Response, error) {
	return requestJSON(http.MethodGet, url, target, nil)
}

func postJSON(url string, target interface{}, payload interface{}) (*http.Response, error) {
	return requestJSON(http.MethodPost, url, target, payload)
}

func deleteJSON(url string, target interface{}) (*http.Response, error) {
	return requestJSON(http.MethodDelete, url, target, nil)
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
