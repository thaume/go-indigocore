// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package cs_test

import (
	"testing"

	"github.com/stratumn/go/cs"
)

func testSegmentValidateError(t *testing.T, s *cs.Segment, want string) {
	if err := s.Validate(); err == nil {
		t.Error("s.Valitate() = nil want Error")
	} else {
		if got := err.Error(); got != want {
			t.Errorf("s.Valitate() = %q want %q", got, want)
		}
	}
}
