// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func cmpWalk(t *testing.T, src, dst, dir string) {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		t.Fatalf("err: filepath.Glob(): %s", err)
	}
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			t.Fatalf("err: os.Stat(): %s", err)
		}
		if info.IsDir() {
			cmpWalk(t, src, dst, file)
			continue
		}
		rel, err := filepath.Rel(src, file)
		if err != nil {
			t.Fatalf("err: filepath.Rel(): %s", err)
		}
		srcPath := filepath.Join(src, rel)
		srcInfo, err := os.Stat(srcPath)
		if err != nil {
			t.Fatalf("err: os.Stat(): %s", err)
		}
		dstPath := filepath.Join(dst, rel)
		dstInfo, err := os.Stat(dstPath)
		if err != nil {
			t.Fatalf("err: os.Stat(): %s", err)
		}
		if got, want := srcInfo.Mode(), dstInfo.Mode(); got != want {
			t.Errorf("err: srcInfo.Mode() = %d want %d", got, want)
		}
		got, err := ioutil.ReadFile(srcPath)
		if err != nil {
			t.Fatalf("err: ioutil.ReadFile(): %s", err)
		}
		want, err := ioutil.ReadFile(dstPath)
		if err != nil {
			t.Fatalf("err: ioutil.ReadFile(): %s", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("err: content =\n%q\nwant\n%q", got, want)
		}
	}
}
