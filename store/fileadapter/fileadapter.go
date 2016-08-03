// A very slow adapter that stores segments as JSON on the file system.
// It is unoptimized (doesn't even use an index) and is designed for development.
package fileadapter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"

	. "github.com/stratumn/go/store/adapter"
	. "github.com/stratumn/go/store/segment"
)

const (
	NAME            = "file"
	DESCRIPTION     = "Stratumn File Adapter"
	DEFAULT_VERSION = "0.1.0"
	DEFAULT_PATH    = "/var/filestore"
)

// The type of the file adapter.
type FileAdapter struct {
	config *Config // adapter config
}

type Config struct {
	Version string
	Path    string // path to directory where files are stored
}

// Creates a new file adapter.
func New(config *Config) *FileAdapter {
	if config.Version == "" {
		config.Version = DEFAULT_VERSION
	}

	return &FileAdapter{config}
}

// Implements github.com/stratumn/go/store/adapter.
func (a *FileAdapter) GetInfo() (interface{}, error) {
	return map[string]interface{}{
		"name":        NAME,
		"description": DESCRIPTION,
		"version":     a.config.Version,
	}, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *FileAdapter) SaveSegment(segment *Segment) error {
	js, err := json.MarshalIndent(segment, "", "  ")
	if err != nil {
		return err
	}

	if err = os.MkdirAll(a.config.Path, 0755); err != nil {
		return err
	}

	segmentPath := path.Join(a.config.Path, segment.Meta["linkHash"].(string)+".json")

	return ioutil.WriteFile(segmentPath, js, 0644)
}

// Implements github.com/stratumn/go/store/adapter.
func (a *FileAdapter) GetSegment(linkHash string) (*Segment, error) {
	file, err := os.Open(path.Join(a.config.Path, linkHash+".json"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var segment Segment

	if err = json.NewDecoder(file).Decode(&segment); err != nil {
		return nil, err
	}

	return &segment, nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *FileAdapter) DeleteSegment(linkHash string) (*Segment, error) {
	segment, err := a.GetSegment(linkHash)
	if segment == nil {
		return segment, err
	}

	if err = os.Remove(path.Join(a.config.Path, linkHash+".json")); err != nil {
		return nil, err
	}

	return segment, err
}

// Implements github.com/stratumn/go/store/adapter.
func (a *FileAdapter) FindSegments(filter *Filter) (SegmentSlice, error) {
	var segments SegmentSlice

	a.forEach(func(segment *Segment) error {
		if filter.MapID != "" && filter.MapID != segment.Link.Meta["mapId"].(string) {
			return nil
		}

		prevLinkHash, ok := segment.Link.Meta["prevLinkHash"].(string)

		if filter.PrevLinkHash != "" && (!ok || filter.PrevLinkHash != prevLinkHash) {
			return nil
		}

		if len(filter.Tags) > 0 {
			if t, ok := segment.Link.Meta["tags"].([]interface{}); ok {
				var tags []string

				for _, v := range t {
					tags = append(tags, v.(string))
				}

				// Exclude the link hash if it doesn't contain all the tags.
				for _, tag := range filter.Tags {
					if !containsString(tags, tag) {
						return nil
					}
				}
			} else {
				return nil
			}
		}

		segments = append(segments, segment)

		return nil
	})

	sort.Sort(segments)

	return paginateSegments(segments, &filter.Pagination), nil
}

// Implements github.com/stratumn/go/store/adapter.
func (a *FileAdapter) GetMapIDs(pagination *Pagination) ([]string, error) {
	set := map[string]bool{}

	a.forEach(func(segment *Segment) error {
		set[segment.Link.Meta["mapId"].(string)] = true
		return nil
	})

	var mapIDs []string

	for mapID := range set {
		mapIDs = append(mapIDs, mapID)
	}

	sort.Strings(mapIDs)

	return paginateStrings(mapIDs, pagination), nil
}

// Loop through all segemnts.
func (a *FileAdapter) forEach(fn func(*Segment) error) error {
	files, err := ioutil.ReadDir(a.config.Path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	re := regexp.MustCompile("(.*)\\.json$")

	for _, file := range files {
		name := file.Name()
		if re.MatchString(name) {
			linkHash := name[:len(name)-5]
			segment, err := a.GetSegment(linkHash)
			if err != nil {
				return err
			}
			if segment == nil {
				return errors.New("could not find segment")
			}
			if err = fn(segment); err != nil {
				return err
			}
		}
	}

	return nil
}

// Returns whether an array contains a string.
func containsString(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

// Paginates a string slice.
func paginateStrings(a []string, p *Pagination) []string {
	length := len(a)

	if p.Offset >= length {
		return []string{}
	}

	if p.Limit > 0 {
		end := min(length, p.Offset+p.Limit)
		return a[p.Offset:end]
	}

	return a[p.Offset:]
}

// Paginates a segment slice.
func paginateSegments(a SegmentSlice, p *Pagination) SegmentSlice {
	length := len(a)

	if p.Offset >= length {
		return SegmentSlice{}
	}

	if p.Limit > 0 {
		end := min(length, p.Offset+p.Limit)
		return a[p.Offset:end]
	}

	return a[p.Offset:]
}

// Min of two ints, duh.
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
