// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// The command dummyfossilizer starts an HTTP server with a dummyfossilizer.
//
// Usage
//
//	$ dummyfossilizer -h
//	Usage of dummyfossilizer:
//	  -callbacktimeout duration
//	    	callback requests timeout (default 10s)
//	  -http string
//	    	http address (default ":6000")
//	  -maxdata int
//	    	maximum data length (default 64)
//	  -mindata int
//	    	minimum data length (default 32)
//	  -tlscert string
//	    	TLS certificate file
//	  -tlskey string
//	    	TLS private key file
//	  -workers int
//	    	number of result workers (default 8)
//
// Docker
//
// A Docker image is available. To create a container, run:
//
//	$ docker run -p 6000:6000 stratumn/dummyfossilizer
package main
