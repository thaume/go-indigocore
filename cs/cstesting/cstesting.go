// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cstesting defines helpers to test Chainscripts.
package cstesting

import (
	"encoding/json"
	"math/rand"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/testutil"
)

// CreateSegment creates a minimal segment.
func CreateSegment(linkHash, mapID, prevLinkHash string, tags []interface{}, priority float64) *cs.Segment {
	linkMeta := map[string]interface{}{
		"mapId":    mapID,
		"priority": priority,
		"random":   testutil.RandomString(12),
	}

	if prevLinkHash != "" {
		linkMeta["prevLinkHash"] = prevLinkHash
	}

	if tags != nil {
		linkMeta["tags"] = tags
	}

	return &cs.Segment{
		Link: cs.Link{
			State: map[string]interface{}{
				"random": testutil.RandomString(12),
			},
			Meta: linkMeta,
		},
		Meta: map[string]interface{}{
			"linkHash": linkHash,
			"random":   testutil.RandomString(12),
		},
	}
}

// RandomSegment creates a random segment.
func RandomSegment() *cs.Segment {
	return CreateSegment(testutil.RandomHash().String(), testutil.RandomString(24),
		testutil.RandomHash().String(), RandomTags(), rand.Float64())
}

// ChangeSegmentState clones a segment and randomly changes its state.
func ChangeSegmentState(s *cs.Segment) *cs.Segment {
	clone := CloneSegment(s)
	clone.Link.State["random"] = testutil.RandomString(12)
	return clone
}

// ChangeSegmentMapID clones a segment and randomly changes its map ID.
func ChangeSegmentMapID(s *cs.Segment) *cs.Segment {
	clone := CloneSegment(s)
	clone.Link.Meta["mapId"] = testutil.RandomString(24)
	return clone
}

// RandomBranch appends a random segment to a segment.
func RandomBranch(s *cs.Segment) *cs.Segment {
	branch := CreateSegment(testutil.RandomHash().String(), testutil.RandomString(24),
		s.Meta["linkHash"].(string), RandomTags(), rand.Float64())
	branch.Link.Meta["mapId"] = s.Link.Meta["mapId"]
	return branch
}

// RandomTags creates between zero and four random tags.
func RandomTags() []interface{} {
	var tags []interface{}
	for i := 0; i < rand.Intn(5); i++ {
		tags = append(tags, testutil.RandomString(12))
	}
	return tags
}

// CloneSegment clones a segment.
func CloneSegment(s *cs.Segment) *cs.Segment {
	var clone cs.Segment

	js, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(js, &clone); err != nil {
		panic(err)
	}

	return &clone
}
