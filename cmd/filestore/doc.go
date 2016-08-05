// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// The command filestore starts an HTTP server with a filestore.
//
// Usage
//
// The following flags are available:
//
//	$ filestore -h
//	Usage of filestore:
//	  -path string
//	    	path to directory where files are stored (default "/var/filestore")
//	  -port string
//	    	server port (default ":5000")
//	  -tlscert string
//	    	TLS certificate file
//	  -tlskey string
//	    	TLS private key file
//	  -verbose
//	    	verbose output
//
// Docker
//
// A Docker image is available. To create a container, run:
//
//	$ docker run -p 5000:5000 -v /path/to/save/files:/var/filestore stratumn/filestore
package main
