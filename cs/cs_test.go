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

func TestSegmentValidate_valid(t *testing.T) {
	s := cstesting.RandomSegment()
	if err := s.Validate(); err != nil {
		t.Errorf("s.Validate() = %q want nil", err)
	}
}

func TestSegmentValidate_linkHashNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Meta, "linkHash")
	testSegmentValidateError(t, s, "meta.linkHash should be a non empty string")
}

func TestSegmentValidate_linkHashEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Meta["linkHash"] = ""
	testSegmentValidateError(t, s, "meta.linkHash should be a non empty string")
}

func TestSegmentValidate_linkHashWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Meta["linkHash"] = 3
	testSegmentValidateError(t, s, "meta.linkHash should be a non empty string")
}

func TestSegmentValidate_mapIDNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "mapId")
	testSegmentValidateError(t, s, "link.meta.mapId should be a non empty string")
}

func TestSegmentValidate_mapIDEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["mapId"] = ""
	testSegmentValidateError(t, s, "link.meta.mapId should be a non empty string")
}

func TestSegmentValidate_mapIDWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["mapId"] = true
	testSegmentValidateError(t, s, "link.meta.mapId should be a non empty string")
}

func TestSegmentValidate_prevLinkHashNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "prevLinkHash")
	if err := s.Validate(); err != nil {
		t.Errorf("s.Validate() = %q want nil", err)
	}
}

func TestSegmentValidate_prevLinkHashEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["prevLinkHash"] = ""
	testSegmentValidateError(t, s, "link.meta.prevLinkHash should be a non empty string")
}

func TestSegmentValidate_prevLinkHashWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["prevLinkHash"] = []string{}
	testSegmentValidateError(t, s, "link.meta.prevLinkHash should be a non empty string")
}

func TestSegmentValidate_tagsNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "tags")
	if err := s.Validate(); err != nil {
		t.Errorf("s.Validate() = %q want nil", err)
	}
}

func TestSegmentValidate_tagsWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = 2.4
	testSegmentValidateError(t, s, "link.meta.tags should be an array of non empty string")
}

func TestSegmentValidate_tagsWrongElementType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = []interface{}{1, true, 3}
	testSegmentValidateError(t, s, "link.meta.tags should be an array of non empty string")
}

func TestSegmentValidate_tagsEmpty(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["tags"] = []interface{}{"test", ""}
	testSegmentValidateError(t, s, "link.meta.tags should be an array of non empty string")
}

func TestSegmentValidate_priorityNil(t *testing.T) {
	s := cstesting.RandomSegment()
	delete(s.Link.Meta, "priority")
	if err := s.Validate(); err != nil {
		t.Errorf("s.Validate() = %q want nil", err)
	}
}

func TestSegmentValidate_priorityWrongType(t *testing.T) {
	s := cstesting.RandomSegment()
	s.Link.Meta["priority"] = false
	testSegmentValidateError(t, s, "link.meta.priority should be a float64")
}

func TestSegmentSliceSort_priority(t *testing.T) {
	slice := cs.SegmentSlice{
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": -1.1}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 3.33}}},
	}

	sort.Sort(slice)
	wantLTE := 100.0
	for i, s := range slice {
		got := s.Link.Meta["priority"].(float64)
		if got > wantLTE {
			t.Errorf("slice#%d: priority = %f want <= %f", i, got, wantLTE)
		}
		wantLTE = got
	}
}

func TestSegmentSliceSort_linkHash(t *testing.T) {
	slice := cs.SegmentSlice{
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.0}}, Meta: map[string]interface{}{"linkHash": "c"}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.0}}, Meta: map[string]interface{}{"linkHash": "b"}},
	}

	sort.Sort(slice)
	wantGTE := "a"
	for i, s := range slice {
		got := s.Meta["linkHash"].(string)
		if got < wantGTE {
			t.Errorf("slice#%d: linkHash = %q want >= %q", i, got, wantGTE)
		}

		wantGTE = got
	}
}

func TestSegmentSliceSort_noPriority(t *testing.T) {
	slice := cs.SegmentSlice{
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{}}},
		&cs.Segment{Link: cs.Link{Meta: map[string]interface{}{"priority": 3.33}}},
	}

	sort.Sort(slice)
	wantLTE := 100.0
	for i, s := range slice {
		got, ok := s.Link.Meta["priority"].(float64)
		if ok {
			if got > wantLTE {
				t.Errorf("slice#%d: priority = %f want <= %f", i, got, wantLTE)
			}

			wantLTE = got
		} else {
			wantLTE = 0
		}
	}
}
