// A store HTTP server with a file adapter.
package main

import (
	"flag"
	"log"

	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/fileadapter"
	"github.com/stratumn/go/store/httpserver"
)

var (
	port     = flag.String("port", httpserver.DEFAULT_PORT, "server port")
	path     = flag.String("path", fileadapter.DEFAULT_PATH, "path to directory where files are stored")
	certFile = flag.String("tlscert", "", "TLS certificate file")
	keyFile  = flag.String("tlskey", "", "TLS private key file")
	verbose  = flag.Bool("verbose", httpserver.DEFAULT_VERBOSE, "verbose output")
	version  = ""
)

func init() {
	log.SetPrefix("filestore ")
}

func main() {
	flag.Parse()

	a := fileadapter.New(&fileadapter.Config{Path: *path, Version: version})
	c := &jsonhttp.Config{
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
		Verbose:  *verbose,
	}
	h := httpserver.New(a, c)

	log.Printf("Listening on %s", *port)
	log.Fatal(h.ListenAndServe())
}
