// Defines a struct for segments.
// Also defines a type for a segment slice and makes it sortable.
package segment

// A type for a segment.
type Segment struct {
	Link Link                   `json:"link"`
	Meta map[string]interface{} `json:"meta"`
}

// A type for a link.
type Link struct {
	State map[string]interface{} `json:"state"`
	Meta  map[string]interface{} `json:"meta"`
}

type SegmentSlice []*Segment

// Implements the sort interface.
func (s SegmentSlice) Len() int {
	return len(s)
}

// Implements the sort interface.
func (s SegmentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Implements the sort interface.
func (s SegmentSlice) Less(i, j int) bool {
	s1 := s[i]
	s2 := s[j]

	p1, ok1 := s1.Link.Meta["priority"].(float64)
	p2, ok2 := s2.Link.Meta["priority"].(float64)

	if !ok1 && ok2 {
		return false
	}

	if ok1 && !ok2 {
		return true
	}

	if ok1 && ok2 {
		return p1 > p2
	}

	return s1.Meta["linkHash"].(string) < s2.Meta["linkHash"].(string)
}
