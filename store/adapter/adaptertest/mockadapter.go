package adaptertest

import (
	. "github.com/stratumn/go/segment"
	. "github.com/stratumn/go/store/adapter"
)

// A type to mock an adapter.
//
// Example usage:
//
// 	adapter := &MockAdapter{}
//	s1 := RandomSegment()
//	adapter.MockGetSegment.Fn = func(linkHash string) (*Segment, error) { return s1, nil }
//	s2, err := adapter.GetSegment("abcdef")     // Should be s1, nil
//	lastArg := adapter.MockGetSegment.LastCalledWith  // Should be "abcdef"
type MockAdapter struct {
	MockGetInfo       MockGetInfo
	MockSaveSegment   MockSaveSegment
	MockGetSegment    MockGetSegment
	MockDeleteSegment MockDeleteSegment
	MockFindSegments  MockFindSegments
	MockGetMapIDs     MockGetMapIDs
}

type MockGetInfo struct {
	CalledCount int
	Fn          func() (interface{}, error)
}

type MockSaveSegment struct {
	CalledCount    int
	CalledWith     []*Segment
	LastCalledWith *Segment
	Fn             func(*Segment) error
}

type MockGetSegment struct {
	CalledCount    int
	CalledWith     []string
	LastCalledWith string
	Fn             func(string) (*Segment, error)
}

type MockDeleteSegment struct {
	CalledCount    int
	CalledWith     []string
	LastCalledWith string
	Fn             func(string) (*Segment, error)
}

type MockFindSegments struct {
	CalledCount    int
	CalledWith     []*Filter
	LastCalledWith *Filter
	Fn             func(*Filter) (SegmentSlice, error)
}

type MockGetMapIDs struct {
	CalledCount    int
	CalledWith     []*Pagination
	LastCalledWith *Pagination
	Fn             func(*Pagination) ([]string, error)
}

// Implements github.com/stratumn/go/store/adapter.
func (a *MockAdapter) GetInfo() (interface{}, error) {
	a.MockGetInfo.CalledCount++

	if a.MockGetInfo.Fn != nil {
		return a.MockGetInfo.Fn()
	}

	return nil, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *MockAdapter) SaveSegment(segment *Segment) error {
	a.MockSaveSegment.CalledCount++
	a.MockSaveSegment.CalledWith = append(a.MockSaveSegment.CalledWith, segment)
	a.MockSaveSegment.LastCalledWith = segment

	if a.MockSaveSegment.Fn != nil {
		return a.MockSaveSegment.Fn(segment)
	}

	return nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *MockAdapter) GetSegment(linkHash string) (*Segment, error) {
	a.MockGetSegment.CalledCount++
	a.MockGetSegment.CalledWith = append(a.MockGetSegment.CalledWith, linkHash)
	a.MockGetSegment.LastCalledWith = linkHash

	if a.MockGetSegment.Fn != nil {
		return a.MockGetSegment.Fn(linkHash)
	}

	return nil, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *MockAdapter) DeleteSegment(linkHash string) (*Segment, error) {
	a.MockDeleteSegment.CalledCount++
	a.MockDeleteSegment.CalledWith = append(a.MockDeleteSegment.CalledWith, linkHash)
	a.MockDeleteSegment.LastCalledWith = linkHash

	if a.MockDeleteSegment.Fn != nil {
		return a.MockDeleteSegment.Fn(linkHash)
	}

	return nil, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *MockAdapter) FindSegments(filter *Filter) (SegmentSlice, error) {
	a.MockFindSegments.CalledCount++
	a.MockFindSegments.CalledWith = append(a.MockFindSegments.CalledWith, filter)
	a.MockFindSegments.LastCalledWith = filter

	if a.MockFindSegments.Fn != nil {
		return a.MockFindSegments.Fn(filter)
	}

	return nil, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *MockAdapter) GetMapIDs(pagination *Pagination) ([]string, error) {
	a.MockGetMapIDs.CalledCount++
	a.MockGetMapIDs.CalledWith = append(a.MockGetMapIDs.CalledWith, pagination)
	a.MockGetMapIDs.LastCalledWith = pagination

	if a.MockGetMapIDs.Fn != nil {
		return a.MockGetMapIDs.Fn(pagination)
	}

	return nil, nil
}
