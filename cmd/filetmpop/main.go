// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"flag"

	"github.com/stratumn/sdk/filestore"
	"github.com/stratumn/sdk/tmpop"
)

var (
	addrPtr = flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	abciPtr = flag.String("tmsp", "socket", "TMSP server: socket | grpc")
	path    = flag.String("path", filestore.DefaultPath, "path to directory where files are stored")
	version = "0.1.0"
	commit  = "00000000000000000000000000000000"
)

func main() {
	flag.Parse()

	a := filestore.New(&filestore.Config{Path: *path, Version: version, Commit: commit})

	tmpop.Run(a, &tmpop.Config{Commit: commit, Version: version, DbDir: *path}, addrPtr, abciPtr)
}
