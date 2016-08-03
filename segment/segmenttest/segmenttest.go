// Provides helpers to test segments.
package segmenttest

import (
	"encoding/json"
	"math/rand"

	. "github.com/stratumn/go/segment"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Generates a random string.
func RandomString(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// Checks if an array contains a string.
func ContainsString(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

// Creates a dummy segment.
func CreateSegment(linkHash, mapId, prevLinkHash string, tags []interface{}, priority float64) *Segment {
	linkMeta := map[string]interface{}{
		"mapId":    mapId,
		"priority": priority,
		"random":   RandomString(12),
	}

	if prevLinkHash != "" {
		linkMeta["prevLinkHash"] = prevLinkHash
	}

	if tags != nil {
		linkMeta["tags"] = tags
	}

	return &Segment{
		Link: Link{
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

// Creates a random dummy segment.
func RandomSegment() *Segment {
	return CreateSegment(RandomString(32), RandomString(24), RandomString(32), RandomTags(), rand.Float64())
}

// Changes the state of a segment, returns a new segment.
func ChangeSegmentState(segment *Segment) *Segment {
	clone := CloneSegment(segment)
	clone.Link.State["random"] = RandomString(12)

	return clone
}

// Changes the map ID of a segment, returns a new segment.
func ChangeSegmentMapID(segment *Segment) *Segment {
	clone := CloneSegment(segment)
	clone.Link.Meta["mapId"] = RandomString(24)

	return clone
}

// Creates a dummy segment branching from a previous segment.
func RandomBranch(segment *Segment) *Segment {
	branch := CreateSegment(RandomString(32), RandomString(24), segment.Meta["linkHash"].(string), RandomTags(), rand.Float64())

	branch.Link.Meta["mapId"] = segment.Link.Meta["mapId"]

	return branch
}

// Creates a random array of tags.
func RandomTags() []interface{} {
	var tags []interface{}

	for i := 0; i < rand.Intn(5); i++ {
		tags = append(tags, RandomString(12))
	}

	return tags
}

// Clones a segment.
func CloneSegment(segment *Segment) *Segment {
	var clone Segment

	js, err := json.Marshal(segment)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(js, &clone); err != nil {
		panic(err)
	}

	return &clone
}
