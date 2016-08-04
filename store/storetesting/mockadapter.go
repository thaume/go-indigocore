package storetesting

import (
	"github.com/stratumn/go/cs"
	"github.com/stratumn/go/store"
)

// MockAdapter is used to mock a store.
//
// It implements github.com/stratumn/go/store.Adapter.
type MockAdapter struct {
	// The mock for the GetInfo function.
	MockGetInfo MockGetInfo

	// The mock for the SaveSegment function.
	MockSaveSegment MockSaveSegment

	// The mock for the GetSegment function.
	MockGetSegment MockGetSegment

	// The mock for the DeleteSegment function.
	MockDeleteSegment MockDeleteSegment

	// The mock for the FindSegments function.
	MockFindSegments MockFindSegments

	// The mock for the GetMapIDs function.
	MockGetMapIDs MockGetMapIDs
}

// MockGetInfo mocks the GetInfo function.
type MockGetInfo struct {
	// The number of times the function was called.
	CalledCount int

	// An optional implementation of the function.
	Fn func() (interface{}, error)
}

// MockSaveSegment mocks the SaveSegment function.
type MockSaveSegment struct {
	// The number of times the function was called.
	CalledCount int

	// The segment that was passed to each call.
	CalledWith []*cs.Segment

	// The last segment that was passed.
	LastCalledWith *cs.Segment

	// An optional implementation of the function.
	Fn func(*cs.Segment) error
}

// MockGetSegment mocks the GetSegment function.
type MockGetSegment struct {
	// The number of times the function was called.
	CalledCount int

	// The link hash that was passed to each call.
	CalledWith []string

	// The last link hash that was passed.
	LastCalledWith string

	// An optional implementation of the function.
	Fn func(string) (*cs.Segment, error)
}

// MockDeleteSegment mocks the DeleteSegment function.
type MockDeleteSegment struct {
	// The number of times the function was called.
	CalledCount int

	// The link hash that was passed to each call.
	CalledWith []string

	// The last link hash that was passed.
	LastCalledWith string

	// An optional implementation of the function.
	Fn func(string) (*cs.Segment, error)
}

// MockFindSegments mocks the FindSegments function.
type MockFindSegments struct {
	// The number of times the function was called.
	CalledCount int

	// The filter that was passed to each call.
	CalledWith []*store.Filter

	// The last filter that was passed.
	LastCalledWith *store.Filter

	// An optional implementation of the function.
	Fn func(*store.Filter) (cs.SegmentSlice, error)
}

// MockGetMapIDs mocks the GetMapIDs function.
type MockGetMapIDs struct {
	// The number of times the function was called.
	CalledCount int

	// The pagination that was passed to each call.
	CalledWith []*store.Pagination

	// The last pagination that was passed.
	LastCalledWith *store.Pagination

	// An optional implementation of the function.
	Fn func(*store.Pagination) ([]string, error)
}

// GetInfo implements github.com/stratumn/go/store.Adapter.GetInfo.
func (a *MockAdapter) GetInfo() (interface{}, error) {
	a.MockGetInfo.CalledCount++

	if a.MockGetInfo.Fn != nil {
		return a.MockGetInfo.Fn()
	}

	return nil, nil
}

// SaveSegment implements github.com/stratumn/go/store.Adapter.SaveSegment.
func (a *MockAdapter) SaveSegment(segment *cs.Segment) error {
	a.MockSaveSegment.CalledCount++
	a.MockSaveSegment.CalledWith = append(a.MockSaveSegment.CalledWith, segment)
	a.MockSaveSegment.LastCalledWith = segment

	if a.MockSaveSegment.Fn != nil {
		return a.MockSaveSegment.Fn(segment)
	}

	return nil
}

// GetSegment implements github.com/stratumn/go/store.Adapter.GetSegment.
func (a *MockAdapter) GetSegment(linkHash string) (*cs.Segment, error) {
	a.MockGetSegment.CalledCount++
	a.MockGetSegment.CalledWith = append(a.MockGetSegment.CalledWith, linkHash)
	a.MockGetSegment.LastCalledWith = linkHash

	if a.MockGetSegment.Fn != nil {
		return a.MockGetSegment.Fn(linkHash)
	}

	return nil, nil
}

// DeleteSegment implements github.com/stratumn/go/store.Adapter.DeleteSegment.
func (a *MockAdapter) DeleteSegment(linkHash string) (*cs.Segment, error) {
	a.MockDeleteSegment.CalledCount++
	a.MockDeleteSegment.CalledWith = append(a.MockDeleteSegment.CalledWith, linkHash)
	a.MockDeleteSegment.LastCalledWith = linkHash

	if a.MockDeleteSegment.Fn != nil {
		return a.MockDeleteSegment.Fn(linkHash)
	}

	return nil, nil
}

// FindSegments implements github.com/stratumn/go/store.Adapter.FindSegments.
func (a *MockAdapter) FindSegments(filter *store.Filter) (cs.SegmentSlice, error) {
	a.MockFindSegments.CalledCount++
	a.MockFindSegments.CalledWith = append(a.MockFindSegments.CalledWith, filter)
	a.MockFindSegments.LastCalledWith = filter

	if a.MockFindSegments.Fn != nil {
		return a.MockFindSegments.Fn(filter)
	}

	return nil, nil
}

// GetMapIDs implements github.com/stratumn/go/store.Adapter.GetMapIDs.
func (a *MockAdapter) GetMapIDs(pagination *store.Pagination) ([]string, error) {
	a.MockGetMapIDs.CalledCount++
	a.MockGetMapIDs.CalledWith = append(a.MockGetMapIDs.CalledWith, pagination)
	a.MockGetMapIDs.LastCalledWith = pagination

	if a.MockGetMapIDs.Fn != nil {
		return a.MockGetMapIDs.Fn(pagination)
	}

	return nil, nil
}
