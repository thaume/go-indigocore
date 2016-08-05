// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// dummyfossilizer starts an HTTP server with a dummyfossilizer.
//
// Usage
//
//	-maxdata int
//	  	maximum data length (default 64)
//	-mindata int
//	  	minimum data length (default 32)
//	-port string
//	  	server port (default ":6000")
//	-tlscert string
//	  	TLS certificate file
//	-tlskey string
//	  	TLS private key file
//	-verbose
//	  	verbose output
//	-workers int
//	  	number of result workers (default 8)
//
// Docker
//
//	docker run -p 6000:6000 stratumn/dummyfossilizer
package main
