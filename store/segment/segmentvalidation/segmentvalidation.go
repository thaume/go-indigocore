// Deals with validating segments.
package segmentvalidation

import (
	"errors"

	. "github.com/stratumn/go/store/segment"
)

// Validates a segment by checking for standard attributes.
func Validate(segment *Segment) error {
	if linkHash, ok := segment.Meta["linkHash"].(string); !ok || linkHash == "" {
		return errors.New("meta.linkHash should be a non empty string")
	}

	if mapId, ok := segment.Link.Meta["mapId"].(string); !ok || mapId == "" {
		return errors.New("link.meta.mapId should be a non empty string")
	}

	if v, ok := segment.Link.Meta["prevLinkHash"]; ok {
		if prevLinkHash, ok := v.(string); !ok || prevLinkHash == "" {
			return errors.New("link.meta.prevLinkHash should be a non empty string")
		}
	}

	if v, ok := segment.Link.Meta["tags"]; ok {
		tags, ok := v.([]interface{})

		if !ok {
			return errors.New("link.meta.tags should be an array of non empty string")
		}

		for _, t := range tags {
			if tag, ok := t.(string); !ok || tag == "" {
				return errors.New("link.meta.tags should be an array of non empty string")
			}
		}
	}

	if v, ok := segment.Link.Meta["priority"]; ok {
		if _, ok := v.(float64); !ok {
			return errors.New("link.meta.priority should be a float64")
		}
	}

	return nil
}
