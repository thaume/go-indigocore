// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package cs_test

import (
	"sort"
	"testing"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/cs/cstesting"
)

func TestSegmentValidateValid(t *testing.T) {
	s := cstesting.RandomSegment()

	if err := s.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestSegmentValidateLinkHashNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Meta, "linkHash")

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateLinkHashEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Meta["linkHash"] = ""

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateLinkHashWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Meta["linkHash"] = 3

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "meta.linkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateMapIDNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "mapId")

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateMapIDEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["mapId"] = ""

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateMapIDWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["mapId"] = true

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.mapId should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidatePrevLinkHashNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "prevLinkHash")

	if err := s.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestSegmentValidatePrevLinkHashEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["prevLinkHash"] = ""

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.prevLinkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidatePrevLinkHashWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["prevLinkHash"] = []string{}

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.prevLinkHash should be a non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateTagsNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "tags")

	if err := s.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestSegmentValidateTagsWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = 2.4

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateTagsWrongElementType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = []interface{}{1, true, 3}

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidateTagsEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = []interface{}{"test", ""}

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.tags should be an array of non empty string" {
		t.Fatal(err)
	}
}

func TestSegmentValidatePriorityNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "priority")

	if err := s.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestSegmentValidatePriorityWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["priority"] = false

	if err := s.Validate(); err == nil {
		t.Fatal("expected error")
	} else if err.Error() != "link.meta.priority should be a float64" {
		t.Fatal(err)
	}
}

func TestSegmentSliceSortable(t *testing.T) {
	slice := cs.SegmentSlice{
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": -1.1}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 3.33}}},
	}

	sort.Sort(slice)

	lastPriority := 100.0

	for _, s := range slice {
		priority := s.Link.Meta["priority"].(float64)
		if priority > lastPriority {
			t.Fatal("expected segments to be sorted by priority")
		}

		lastPriority = priority
	}
}

func TestSegmentSliceSortableLinkHash(t *testing.T) {
	slice := cs.SegmentSlice{
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.0}}, Meta: map[string]interface{}{"linkHash": "c"}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.0}}, Meta: map[string]interface{}{"linkHash": "b"}},
	}

	sort.Sort(slice)

	lastLinkHash := "a"

	for _, s := range slice {
		linkHash := s.Meta["linkHash"].(string)
		if linkHash < lastLinkHash {
			t.Fatal("expected segments to be sorted by link hashes")
		}

		lastLinkHash = linkHash
	}
}

func TestSegmentSliceSortableNoPriority(t *testing.T) {
	slice := cs.SegmentSlice{
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 3.33}}},
	}

	sort.Sort(slice)

	lastPriority := 100.0

	for _, s := range slice {
		priority, ok := s.Link.Meta["priority"].(float64)
		if ok {
			if priority > lastPriority {
				t.Fatal("expected segments to be sorted by priority")
			}

			lastPriority = priority
		} else {
			lastPriority = 0
		}
	}
}
