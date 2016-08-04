// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

// A store HTTP server with a file adapter.
package main

import (
	"flag"
	"log"

	"github.com/stratumn/go/filestore"
	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/storehttp"
)

var (
	port     = flag.String("port", storehttp.DefaultPort, "server port")
	path     = flag.String("path", filestore.DefaultPath, "path to directory where files are stored")
	certFile = flag.String("tlscert", "", "TLS certificate file")
	keyFile  = flag.String("tlskey", "", "TLS private key file")
	verbose  = flag.Bool("verbose", storehttp.DefaultVerbose, "verbose output")
	version  = ""
)

func init() {
	log.SetPrefix("filestore ")
}

func main() {
	flag.Parse()

	a := filestore.New(&filestore.Config{Path: *path, Version: version})
	c := &jsonhttp.Config{
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
		Verbose:  *verbose,
	}
	h := storehttp.New(a, c)

	log.Printf("Listening on %s", *port)
	log.Fatal(h.ListenAndServe())
}
