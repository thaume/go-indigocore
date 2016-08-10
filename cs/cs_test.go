// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package cs

import (
	"sort"
	"testing"
)

func TestSortable(t *testing.T) {
	slice := SegmentSlice{
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": -1.1}}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 3.33}}},
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

func TestSortableLinkHash(t *testing.T) {
	slice := SegmentSlice{
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 2.0}}, Meta: map[string]interface{}{"linkHash": "c"}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 2.0}}, Meta: map[string]interface{}{"linkHash": "b"}},
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

func TestSortableNoPriority(t *testing.T) {
	slice := SegmentSlice{
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&Segment{Link: Link{Meta: map[string]interface{}{}}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 3.33}}},
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
