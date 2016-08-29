# Stratumn store

A Golang package to create Stratumn stores.

## Adapters

An adapter must implement this interface:

```go
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
```

You can then use `github.com/stratumn/go/store/storehttp` to create an HTTP server for that adapter.

See `github.com/stratumn/go/filestore` or `github.com/stratumn/go/dummystore` for an example.
