package segmentvalidation

import (
	"testing"

	. "github.com/stratumn/go/store/segment/segmenttest"
)

func TestValidateValid(t *testing.T) {
	segment := RandomSegment()

	if err := Validate(segment); err != nil {
		t.Fatal(err)
	}
}

func TestValidateLinkHashNil(t *testing.T) {
	segment := RandomSegment()
	delete(segment.Meta, "linkHash")

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateLinkHashEmpty(t *testing.T) {
	segment := RandomSegment()
	segment.Meta["linkHash"] = ""

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateLinkHashWrongType(t *testing.T) {
	segment := RandomSegment()
	segment.Meta["linkHash"] = 3

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateMapIDNil(t *testing.T) {
	segment := RandomSegment()
	delete(segment.Link.Meta, "mapId")

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateMapIDEmpty(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["mapId"] = ""

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateMapIDWrongType(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["mapId"] = true

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidatePrevLinkHashNil(t *testing.T) {
	segment := RandomSegment()
	delete(segment.Link.Meta, "prevLinkHash")

	if err := Validate(segment); err != nil {
		t.Fatal(err)
	}
}

func TestValidatePrevLinkHashEmpty(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["prevLinkHash"] = ""

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.prevLinkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidatePrevLinkHashWrongType(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["prevLinkHash"] = []string{}

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.prevLinkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateTagsNil(t *testing.T) {
	segment := RandomSegment()
	delete(segment.Link.Meta, "tags")

	if err := Validate(segment); err != nil {
		t.Fatal(err)
	}
}

func TestValidateTagsWrongType(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["tags"] = 2.4

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestValidateTagsWrongElementType(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["tags"] = []int{1, 2, 3}

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestValidatePriorityNil(t *testing.T) {
	segment := RandomSegment()
	delete(segment.Link.Meta, "priority")

	if err := Validate(segment); err != nil {
		t.Fatal(err)
	}
}

func TestValidatePriorityWrongType(t *testing.T) {
	segment := RandomSegment()
	segment.Link.Meta["priority"] = false

	if err := Validate(segment); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.priority should be a float64" {
		t.Fatal(err)
	}
}
