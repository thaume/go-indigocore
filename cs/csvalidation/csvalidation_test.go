// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package csvalidation

import (
	"testing"

	"github.com/stratumn/go/cs/cstesting"
)

func TestValidateValid(t *testing.T) {
	s := cstesting.RandomSegment()

	if err := Validate(s); err != nil {
		t.Fatal(err)
	}
}

func TestValidateLinkHashNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Meta, "linkHash")

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateLinkHashEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Meta["linkHash"] = ""

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateLinkHashWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Meta["linkHash"] = 3

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateMapIDNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "mapId")

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateMapIDEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["mapId"] = ""

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateMapIDWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["mapId"] = true

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidatePrevLinkHashNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "prevLinkHash")

	if err := Validate(s); err != nil {
		t.Fatal(err)
	}
}

func TestValidatePrevLinkHashEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["prevLinkHash"] = ""

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.prevLinkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidatePrevLinkHashWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["prevLinkHash"] = []string{}

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.prevLinkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestValidateTagsNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "tags")

	if err := Validate(s); err != nil {
		t.Fatal(err)
	}
}

func TestValidateTagsWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = 2.4

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestValidateTagsWrongElementType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = []interface{}{1, true, 3}

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestValidateTagsEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = []interface{}{"test", ""}

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestValidatePriorityNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "priority")

	if err := Validate(s); err != nil {
		t.Fatal(err)
	}
}

func TestValidatePriorityWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["priority"] = false

	if err := Validate(s); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.priority should be a float64" {
		t.Fatal(err)
	}
}
