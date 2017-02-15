// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package jsonhttp

import "testing"

func testErrStatus(t *testing.T, err ErrHTTP, want int) {
	if got := err.Status(); got != want {
		t.Errorf("err.Status() = %d want %d", got, want)
	}
}

func testErrError(t *testing.T, err ErrHTTP, want string) {
	if got := err.Error(); got != want {
		t.Errorf("err.Error() = %q want %q", got, want)
	}
}
