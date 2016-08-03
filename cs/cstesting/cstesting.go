// Package cstesting defines helpers to test Chainscripts.
package cstesting

import (
	"encoding/json"
	"math/rand"

	"github.com/stratumn/go/cs"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString generates a random string.
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ContainsString checks if an array contains a string.
func ContainsString(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// CreateSegment creates a minimal segment.
func CreateSegment(linkHash, mapID, prevLinkHash string, tags []interface{}, priority float64) *cs.Segment {
	linkMeta := map[string]interface{}{
		"mapId":    mapID,
		"priority": priority,
		"random":   RandomString(12),
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
				"random": RandomString(12),
			},
			Meta: linkMeta,
		},
		Meta: map[string]interface{}{
			"linkHash": linkHash,
			"random":   RandomString(12),
		},
	}
}

// RandomSegment creates a random segment.
func RandomSegment() *cs.Segment {
	return CreateSegment(RandomString(32), RandomString(24), RandomString(32), RandomTags(), rand.Float64())
}

// ChangeSegmentState clones a segment and randomly changes its state.
func ChangeSegmentState(s *cs.Segment) *cs.Segment {
	clone := CloneSegment(s)
	clone.Link.State["random"] = RandomString(12)
	return clone
}

// ChangeSegmentMapID clones a segment and randomly changes its map ID.
func ChangeSegmentMapID(s *cs.Segment) *cs.Segment {
	clone := CloneSegment(s)
	clone.Link.Meta["mapId"] = RandomString(24)
	return clone
}

// RandomBranch appends a random segment to a segment.
func RandomBranch(s *cs.Segment) *cs.Segment {
	branch := CreateSegment(RandomString(32), RandomString(24), s.Meta["linkHash"].(string), RandomTags(), rand.Float64())
	branch.Link.Meta["mapId"] = s.Link.Meta["mapId"]
	return branch
}

// RandomTags creates between zero and four random tags.
func RandomTags() []interface{} {
	var tags []interface{}
	for i := 0; i < rand.Intn(5); i++ {
		tags = append(tags, RandomString(12))
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
