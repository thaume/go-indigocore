// Copyright 2016 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cs defines types to work with Chainscripts.
package cs

import (
	"errors"
)

// Segment contains a link and meta data about the link.
type Segment struct {
	Link Link                   `json:"link"`
	Meta map[string]interface{} `json:"meta"`
}

// Validate checks for errors in a segment.
func (s *Segment) Validate() error {
	if linkHash, ok := s.Meta["linkHash"].(string); !ok || linkHash == "" {
		return errors.New("meta.linkHash should be a non empty string")
	}
	if mapID, ok := s.Link.Meta["mapId"].(string); !ok || mapID == "" {
		return errors.New("link.meta.mapId should be a non empty string")
	}
	if v, ok := s.Link.Meta["prevLinkHash"]; ok {
		if prevLinkHash, ok := v.(string); !ok || prevLinkHash == "" {
			return errors.New("link.meta.prevLinkHash should be a non empty string")
		}
	}

	if v, ok := s.Link.Meta["tags"]; ok {
		tags, ok := v.([]interface{})
		if !ok {
			return errors.New("link.meta.tags should be an array of non empty string")
		}
		for _, t := range tags {
			if tag, ok := t.(string); !ok || tag == "" {
				return errors.New("link.meta.tags should be an array of non empty string")
			}
		}
	}

	if v, ok := s.Link.Meta["priority"]; ok {
		if _, ok := v.(float64); !ok {
			return errors.New("link.meta.priority should be a float64")
		}
	}

	return nil
}

// Link contains a state and meta data about the state.
type Link struct {
	State map[string]interface{} `json:"state"`
	Meta  map[string]interface{} `json:"meta"`
}

// SegmentSlice is a slice of segment pointers.
type SegmentSlice []*Segment

// Len implements sort.Interface.Len.
func (s SegmentSlice) Len() int {
	return len(s)
}

// Swap implements sort.Interface.Swap.
func (s SegmentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less implements sort.Interface.Less.
func (s SegmentSlice) Less(i, j int) bool {
	var (
		s1      = s[i]
		s2      = s[j]
		p1, ok1 = s1.Link.Meta["priority"].(float64)
		p2, ok2 = s2.Link.Meta["priority"].(float64)
	)

	if !ok1 && ok2 {
		return false
	}

	if ok1 && !ok2 {
		return true
	}

	if ok1 && ok2 && p1 != p2 {
		return p1 > p2
	}

	return s1.Meta["linkHash"].(string) < s2.Meta["linkHash"].(string)
}
