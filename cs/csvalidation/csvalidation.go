// Package csvalidation defines functionality to validate Chainscripts.
package csvalidation

import (
	"errors"

	"github.com/stratumn/go/cs"
)

// Validate checks for errors in a segment.
func Validate(s *cs.Segment) error {
	if linkHash, ok := s.Meta["linkHash"].(string); !ok || linkHash == "" {
		return errors.New("meta.linkHash should be a non empty string")
	}

	if mapID, ok := s.Link.Meta["mapId"].(string); !ok || mapID == "" {
		return errors.New("link.meta.mapId should be a non empty string")
	}

	if v, ok := s.Link.Meta["prevLinkHash"]; ok {
		if prevLinkHash, ok := v.(string); !ok || prevLinkHash == "" {
			return errors.New("link.meta.prevLinkHash should be a non empty string")
		}
	}

	if v, ok := s.Link.Meta["tags"]; ok {
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

	if v, ok := s.Link.Meta["priority"]; ok {
		if _, ok := v.(float64); !ok {
			return errors.New("link.meta.priority should be a float64")
		}
	}

	return nil
}
