package adaptertest

import (
	"testing"

	. "github.com/stratumn/go/store/adapter"
)

// Tests what happens when you get information about the adapter.
func TestGetInfo(t *testing.T, adapter Adapter) {
	info, err := adapter.GetInfo()

	if info == nil {
		t.Fatal("info is nil")
	}

	if err != nil {
		t.Fatal(err)
	}
}
