// A naive adapter that stores segments in memory.
// It is unoptimized and is designed for testing.
package dummyadapter

import (
	"sort"
	"sync"

	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/store/adapter"
)

const (
	NAME            = "dummy"
	DESCRIPTION     = "Stratumn Dummy Adapter"
	DEFAULT_VERSION = "0.1.0"
)

// The type of the dummy adapter.
type DummyAdapter struct {
	version  string
	segments segmentMap    // maps link hashes to segments
	maps     hashSetMap    // maps chains IDs to sets of link hashes
	mutex    *sync.RWMutex // simple global mutex, just in case
}

type segmentMap map[string]*Segment
type hashSet map[string]bool
type hashSetMap map[string]hashSet

// Creates a new dummy adapter.
func New(version string) *DummyAdapter {
	if version == "" {
		version = DEFAULT_VERSION
	}

	return &DummyAdapter{version, segmentMap{}, hashSetMap{}, &sync.RWMutex{}}
}

// Implements github.com/stratumn/go/store/adapter.
func (a *DummyAdapter) GetInfo() (interface{}, error) {
	return map[string]interface{}{
		"name":        NAME,
		"description": DESCRIPTION,
		"version":     a.version,
	}, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *DummyAdapter) SaveSegment(segment *Segment) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	linkHash := segment.Meta["linkHash"].(string)

	curr, err := a.GetSegment(linkHash)

	if err != nil {
		return err
	}

	mapID := segment.Link.Meta["mapId"].(string)

	if curr != nil {
		// Remove current segment from map if needed
		currMapId := curr.Link.Meta["mapId"].(string)

		if currMapId != mapID {
			delete(a.maps[currMapId], linkHash)
		}
	}

	_, exists := a.maps[mapID]

	if !exists {
		a.maps[mapID] = hashSet{}
	}

	a.segments[linkHash] = segment
	a.maps[mapID][linkHash] = true

	return nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *DummyAdapter) GetSegment(linkHash string) (*Segment, error) {
	return a.segments[linkHash], nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *DummyAdapter) DeleteSegment(linkHash string) (*Segment, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	segment, exists := a.segments[linkHash]

	if !exists {
		return nil, nil
	}

	delete(a.segments, linkHash)
	delete(a.maps[segment.Link.Meta["mapId"].(string)], linkHash)

	return segment, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *DummyAdapter) FindSegments(filter *Filter) (SegmentSlice, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	var linkHashes hashSet
	var exists bool

	if filter.MapID == "" {
		linkHashes = hashSet{}

		for linkHash := range a.segments {
			linkHashes[linkHash] = true
		}
	} else {
		linkHashes, exists = a.maps[filter.MapID]

		if !exists {
			return SegmentSlice{}, nil
		}
	}

	return a.findHashesSegments(linkHashes, filter)
}

// Implements github.com/stratumn/go/store/adapter.
func (a *DummyAdapter) GetMapIDs(pagination *Pagination) ([]string, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	mapIDs := make([]string, len(a.maps))
	i := 0

	for mapID := range a.maps {
		mapIDs[i] = mapID
		i++
	}

	sort.Strings(mapIDs)

	return paginateStrings(mapIDs, pagination), nil
}

// Given a set of link hashes and a filter, returns a slice of segments that match.
func (a *DummyAdapter) findHashesSegments(linkHashes hashSet, filter *Filter) (SegmentSlice, error) {
	var segments SegmentSlice

HASH_LOOP:
	for linkHash := range linkHashes {
		segment := a.segments[linkHash]

		prevLinkHash, ok := segment.Link.Meta["prevLinkHash"].(string)

		if filter.PrevLinkHash != "" && (!ok || filter.PrevLinkHash != prevLinkHash) {
			continue
		}

		if len(filter.Tags) > 0 {
			if t, ok := segment.Link.Meta["tags"].([]interface{}); ok {
				var tags []string

				for _, v := range t {
					tags = append(tags, v.(string))
				}

				// Exclude the link hash if it doesn't contain all the tags.
				for _, tag := range filter.Tags {
					if !containsString(tags, tag) {
						continue HASH_LOOP
					}
				}
			} else {
				continue HASH_LOOP
			}
		}

		segments = append(segments, segment)
	}

	// SegmentSlice implements the sort interface to sort by priority.
	sort.Sort(segments)

	return paginateSegments(segments, &filter.Pagination), nil
}

// Returns whether an array contains a string.
func containsString(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

// Paginates a string slice.
func paginateStrings(a []string, p *Pagination) []string {
	length := len(a)

	if p.Offset >= length {
		return []string{}
	}

	if p.Limit > 0 {
		end := min(length, p.Offset+p.Limit)
		return a[p.Offset:end]
	}

	return a[p.Offset:]
}

// Paginates a segment slice.
func paginateSegments(a SegmentSlice, p *Pagination) SegmentSlice {
	length := len(a)

	if p.Offset >= length {
		return SegmentSlice{}
	}

	if p.Limit > 0 {
		end := min(length, p.Offset+p.Limit)
		return a[p.Offset:end]
	}

	return a[p.Offset:]
}

// Min of two ints, duh.
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
