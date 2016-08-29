// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package jsonhttp

import "testing"

func testErrStatus(t *testing.T, err ErrHTTP, want int) {
	if got := err.Status(); want != got {
		t.Errorf("err.Status() = %d want %d", got, want)
	}
}

func testErrError(t *testing.T, err ErrHTTP, want string) {
	if got := err.Error(); want != got {
		t.Errorf("err.Error() = %q want %q", got, want)
	}
}
