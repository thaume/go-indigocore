package storetesting

import (
	"testing"

	"github.com/stratumn/go/store"
)

// TestGetInfo tests what happens when you get information about the adapter.
func TestGetInfo(t *testing.T, a store.Adapter) {
	info, err := a.GetInfo()

	if info == nil {
		t.Fatal("info is nil")
	}

	if err != nil {
		t.Fatal(err)
	}
}
