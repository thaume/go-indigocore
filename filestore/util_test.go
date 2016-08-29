// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package filestore

import (
	"io/ioutil"
	"testing"
)

func createAdapter(tb testing.TB) *FileStore {
	path, err := ioutil.TempDir("", "filestore")
	if err != nil {
		tb.Fatalf("ioutil.TempDir(): err: %s", err)
	}
	return New(&Config{Path: path})
}
