// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// Package dummystore implements a store that saves all the segments in memory.
//
// It can be used for testing, but it's unoptimized and not designed for production.
package dummystore

import (
	"sort"
	"sync"

	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/store"
	"github.com/stratumn/go/types"
)

const (
	// Name is the name set in the store's information.
	Name = "dummy"

	// Description is the description set in the store's information.
	Description = "Stratumn Dummy Store"
)

// Config contains configuration options for the store.
type Config struct {
	// A version string that will set in the store's information.
	Version string

	// A git commit hash that will set in the store's information.
	Commit string
}

// Info is the info returned by GetInfo.
type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Commit      string `json:"commit"`
}

// DummyStore is the type that implements github.com/stratumn/go/store.Adapter.
type DummyStore struct {
	config   *Config
	segments segmentMap   // maps link hashes to segments
	maps     hashSetMap   // maps chains IDs to sets of link hashes
	mutex    sync.RWMutex // simple global mutex, just in case
}

type segmentMap map[string]*cs.Segment
type hashSet map[string]struct{}
type hashSetMap map[string]hashSet

// New creates an instance of a DummyStore.
func New(config *Config) *DummyStore {
	return &DummyStore{config, segmentMap{}, hashSetMap{}, sync.RWMutex{}}
}

// GetInfo implements github.com/stratumn/go/store.Adapter.GetInfo.
func (a *DummyStore) GetInfo() (interface{}, error) {
	return &Info{
		Name:        Name,
		Description: Description,
		Version:     a.config.Version,
		Commit:      a.config.Commit,
	}, nil
}

// SaveSegment implements github.com/stratumn/go/store.Adapter.SaveSegment.
func (a *DummyStore) SaveSegment(segment *cs.Segment) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	linkHashStr := segment.Meta["linkHash"].(string)
	linkHash, err := types.NewBytes32FromString(linkHashStr)
	if err != nil {
		return err
	}

	curr, err := a.GetSegment(linkHash)
	if err != nil {
		return err
	}

	mapID := segment.Link.Meta["mapId"].(string)

	if curr != nil {
		currMapID := curr.Link.Meta["mapId"].(string)
		if currMapID != mapID {
			delete(a.maps[currMapID], linkHashStr)
		}
	}

	_, exists := a.maps[mapID]
	if !exists {
		a.maps[mapID] = hashSet{}
	}

	a.segments[linkHashStr] = segment
	a.maps[mapID][linkHashStr] = struct{}{}

	return nil
}

// GetSegment implements github.com/stratumn/go/store.Adapter.GetSegment.
func (a *DummyStore) GetSegment(linkHash *types.Bytes32) (*cs.Segment, error) {
	return a.segments[linkHash.String()], nil
}

// DeleteSegment implements github.com/stratumn/go/store.Adapter.DeleteSegment.
func (a *DummyStore) DeleteSegment(linkHash *types.Bytes32) (*cs.Segment, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	linkHashStr := linkHash.String()
	segment, exists := a.segments[linkHashStr]
	if !exists {
		return nil, nil
	}

	delete(a.segments, linkHashStr)
	delete(a.maps[segment.Link.Meta["mapId"].(string)], linkHashStr)

	return segment, nil
}

// FindSegments implements github.com/stratumn/go/store.Adapter.FindSegments.
func (a *DummyStore) FindSegments(filter *store.Filter) (cs.SegmentSlice, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	var (
		linkHashes hashSet
		exists     bool
	)

	if filter.MapID == "" {
		linkHashes = hashSet{}
		for linkHash := range a.segments {
			linkHashes[linkHash] = struct{}{}
		}
	} else {
		linkHashes, exists = a.maps[filter.MapID]
		if !exists {
			return cs.SegmentSlice{}, nil
		}
	}

	return a.findHashesSegments(linkHashes, filter)
}

// GetMapIDs implements github.com/stratumn/go/store.Adapter.GetMapIDs.
func (a *DummyStore) GetMapIDs(pagination *store.Pagination) ([]string, error) {
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

func (a *DummyStore) findHashesSegments(linkHashes hashSet, filter *store.Filter) (cs.SegmentSlice, error) {
	var segments cs.SegmentSlice

HASH_LOOP:
	for linkHash := range linkHashes {
		segment := a.segments[linkHash]

		if filter.PrevLinkHash != nil {
			prevLinkHash, ok := segment.Link.Meta["prevLinkHash"].(string)
			if !ok || filter.PrevLinkHash.String() != prevLinkHash {
				continue
			}
		}

		if len(filter.Tags) > 0 {
			if t, ok := segment.Link.Meta["tags"].([]interface{}); ok {
				var tags []string
				for _, v := range t {
					tags = append(tags, v.(string))
				}
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

	sort.Sort(segments)

	return paginateSegments(segments, &filter.Pagination), nil
}

func containsString(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func paginateStrings(a []string, p *store.Pagination) []string {
	l := len(a)
	if p.Offset >= l {
		return []string{}
	}

	if p.Limit > 0 {
		end := min(l, p.Offset+p.Limit)
		return a[p.Offset:end]
	}

	return a[p.Offset:]
}

func paginateSegments(a cs.SegmentSlice, p *store.Pagination) cs.SegmentSlice {
	l := len(a)
	if p.Offset >= l {
		return cs.SegmentSlice{}
	}

	if p.Limit > 0 {
		end := min(l, p.Offset+p.Limit)
		return a[p.Offset:end]
	}

	return a[p.Offset:]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
