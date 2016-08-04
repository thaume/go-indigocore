package storehttp_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/stratumn/go/dummystore"
	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/storehttp"
)

func Example() {
	// Create a dummy adapter.
	a := dummystore.New("0.1.0")
	c := &jsonhttp.Config{
		Port: "5555",
	}

	// Create a server.
	s := storehttp.New(a, c)

	// Create a test server.
	ts := httptest.NewServer(s)
	defer ts.Close()

	// Test the root route.
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	info, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", info)
	// Output: {"adapter":{"description":"Stratumn Dummy Store","name":"dummy","version":"0.1.0"}}
}
