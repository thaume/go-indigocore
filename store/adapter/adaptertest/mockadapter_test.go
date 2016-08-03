package adaptertest

import (
	"reflect"
	"testing"

	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/segment/segmenttest"
)

func TestMockAdapter(t *testing.T) {
	adapter := &MockAdapter{}
	segment1 := RandomSegment()
	adapter.MockGetSegment.Fn = func(linkHash string) (*Segment, error) { return segment1, nil }
	segment2, err := adapter.GetSegment("abcdef")

	if segment1 != segment2 {
		t.Fatal("expected segments to be equal")
	}

	if err != nil {
		t.Fatal("unexpected error")
	}

	adapter.GetSegment("ghij")

	if adapter.MockGetSegment.CalledCount != 2 {
		t.Fatal("unexpected MockGetSegment.CalledCount value")
	}

	if !reflect.DeepEqual(adapter.MockGetSegment.CalledWith, []string{"abcdef", "ghij"}) {
		t.Fatal("unexpected MockGetSegment.LastCalledWith value")
	}

	if adapter.MockGetSegment.LastCalledWith != "ghij" {
		t.Fatal("unexpected MockGetSegment.LastCalledWith value")
	}
}
