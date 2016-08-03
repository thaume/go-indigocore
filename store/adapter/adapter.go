// Defines an interface to create an adapter.
package adapter

import (
	. "github.com/stratumn/go/store/segment"
)

// An adapter handles creating and fetching segments.
type Adapter interface {
	// Returns arbitrary information about the adapter.
	GetInfo() (interface{}, error)

	// Creates or updates a segment. Segments passed to this method are
	// assumed to be valid.
	SaveSegment(segment *Segment) error

	// Get a segment by link hash. Returns nil if no match is found.
	GetSegment(linkHash string) (*Segment, error)

	// Deletes a segment by link hash. Returns the removed segment or nil
	// if not found.
	DeleteSegment(linkHash string) (*Segment, error)

	// Find segments. Returns an empty slice if there are no results.
	FindSegments(filter *Filter) (SegmentSlice, error)

	// Get all the existing map IDs.
	GetMapIDs(pagination *Pagination) ([]string, error)
}

// Pagination options.
type Pagination struct {
	Offset int // index of the first segment
	Limit  int // maximum number of segments, all if zero
}

// Filtering options.
type Filter struct {
	Pagination
	MapID        string   // a map ID the segments must have
	PrevLinkHash string   // a previous link hash the segments must have
	Tags         []string // a slice of tags the segments must contains
}
