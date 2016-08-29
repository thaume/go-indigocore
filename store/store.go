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

// Package store defines types to implement a store.
package store

import (
	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/types"
)

// Adapter must be implemented by a store.
type Adapter interface {
	// Returns arbitrary information about the adapter.
	GetInfo() (interface{}, error)

	// Creates or updates a segment. Segments passed to this method are
	// assumed to be valid.
	SaveSegment(segment *cs.Segment) error

	// Get a segment by link hash. Returns nil if no match is found.
	GetSegment(linkHash *types.Bytes32) (*cs.Segment, error)

	// Deletes a segment by link hash. Returns the removed segment or nil
	// if not found.
	DeleteSegment(linkHash *types.Bytes32) (*cs.Segment, error)

	// Find segments. Returns an empty slice if there are no results.
	FindSegments(filter *Filter) (cs.SegmentSlice, error)

	// Get all the existing map IDs.
	GetMapIDs(pagination *Pagination) ([]string, error)
}

// Pagination contains pagination options.
type Pagination struct {
	// Index of the first segment.
	Offset int

	// Maximum number of segments, all if zero.
	Limit int
}

// Filter contains filtering options.
type Filter struct {
	Pagination

	// A map ID the segments must have.
	MapID string

	// A previous link hash the segments must have.
	PrevLinkHash *types.Bytes32

	// A slice of tags the segments must contains.
	Tags []string
}
