package dummyadapter

import (
	"testing"

	. "github.com/stratumn/go/fossilizer/adapter"
)

func TestFossilize(t *testing.T) {
	adapter := New("")

	resultChan := make(chan *Result)

	adapter.AddResultChan(resultChan)

	data := []byte("data")
	meta := []byte("meta")

	go func() {
		if err := adapter.Fossilize(data, meta); err != nil {
			t.Fatal(err)
		}
	}()

	result := <-resultChan

	if string(result.Data) != string(data) {
		t.Fatal("Unexpected result data")
	}

	if string(result.Meta) != string(meta) {
		t.Fatal("Unexpected result meta")
	}

	if result.Evidence.(map[string]interface{})["authority"].(string) != "dummy" {
		t.Fatal("Unexpected result evidence")
	}
}
