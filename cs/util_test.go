// Copyright 2016 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
