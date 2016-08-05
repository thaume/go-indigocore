// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// The command dummystore starts an HTTP server with a dummystore.
//
// Usage
//
// The following flags are available:
//
//	$ dummystore -h
//	Usage of dummystore:
//	  -port string
//	  	server port (default ":5000")
//	  -tlscert string
//	  	TLS certificate file
//	  -tlskey string
//	  	TLS private key file
//	  -verbose
//	    	verbose output
//
// Docker
//
// A Docker image is available. To create a container, run:
//
//	$ docker run -p 5000:5000 stratumn/dummystore
package main
