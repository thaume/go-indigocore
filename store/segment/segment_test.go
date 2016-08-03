package segment

import (
	"sort"
	"testing"
)

func TestSortable(t *testing.T) {
	slice := SegmentSlice{
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": -1.1}}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 3.33}}},
	}

	sort.Sort(slice)

	lastPriority := 100.0

	for _, segment := range slice {
		priority := segment.Link.Meta["priority"].(float64)

		if priority > lastPriority {
			t.Fatal("expected segments to be sorted by priority")
		}

		lastPriority = priority
	}
}

func TestSortableNoPriority(t *testing.T) {
	slice := SegmentSlice{
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 2.3}}},
		&Segment{Link: Link{Meta: map[string]interface{}{}}},
		&Segment{Link: Link{Meta: map[string]interface{}{"priority": 3.33}}},
	}

	sort.Sort(slice)

	lastPriority := 100.0

	for _, segment := range slice {
		priority, ok := segment.Link.Meta["priority"].(float64)

		if ok {
			if priority > lastPriority {
				t.Fatal("expected segments to be sorted by priority")
			}

			lastPriority = priority
		} else {
			lastPriority = 0
		}
	}
}
