package dummyadapter

import (
	"testing"

	. "github.com/stratumn/go/fossilizer/adapter"
)

func TestFossilize(t *testing.T) {
	a := New("")

	resultChan := make(chan *Result)

	a.AddResultChan(resultChan)

	data := []byte("data")
	meta := []byte("meta")

	go func() {
		if err := a.Fossilize(data, meta); err != nil {
			t.Fatal(err)
		}
	}()

	r := <-resultChan

	if string(r.Data) != string(data) {
		t.Fatal("Unexpected result data")
	}

	if string(r.Meta) != string(meta) {
		t.Fatal("Unexpected result meta")
	}

	if r.Evidence.(map[string]interface{})["authority"].(string) != "dummy" {
		t.Fatal("Unexpected result evidence")
	}
}
