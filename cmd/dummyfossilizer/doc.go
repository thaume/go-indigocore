// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// The command dummyfossilizer starts an HTTP server with a dummyfossilizer.
//
// Usage
//
// The following flags are available:
//
//	$ dummyfossilizer -h
//	Usage of dummyfossilizer:
//	  -callbacktimeout duration
//	    	callback requests timeout (default 10s)
//	  -maxdata int
//	    	maximum data length (default 64)
//	  -mindata int
//	    	minimum data length (default 32)
//	  -port string
//	    	server port (default ":6000")
//	  -tlscert string
//	    	TLS certificate file
//	  -tlskey string
//	    	TLS private key file
//	  -verbose
//	    	verbose output
//	  -workers int
//	    	number of result workers (default 8)
//
// Docker
//
// A Docker image is available. To create a container, run:
//
//	$ docker run -p 6000:6000 stratumn/dummyfossilizer
package main
