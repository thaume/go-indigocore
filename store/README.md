# Stratumn store

A Golang package to create Stratumn stores.

## Adapters

An adapter must implement this interface:

```go
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
```

You can then use `github.com/stratumn/go/store/httpserver` to create an HTTP server for that adapter.

See `github.com/stratumn/go/store/fileadapter` or `github.com/stratumn/go/store/dummyadapter` for an example.
