// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// The command filetmpop starts a server.
//
// Usage:
//
//
//	$ filetmpop -h
// 		Usage of dist/darwin-amd64/filetmpop:
// 		  -addr string
// 			Listen address (default "tcp://0.0.0.0:46658")
// 		  -path string
// 			path to directory where files are stored (default "/var/stratumn/filestore")
// 		  -tmsp string
// 			TMSP server: socket | grpc (default "socket")
//
// Docker:
//
// A Docker image is available. To create a container, run:
//
//	$ docker run -p 46658:46658 stratumn/filetmpop

package main
