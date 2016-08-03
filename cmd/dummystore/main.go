// A store HTTP server with a dummy adapter.
package main

import (
	"flag"
	"log"

	"github.com/stratumn/go/store/dummyadapter"
	"github.com/stratumn/go/store/httpserver"
)

var (
	port     = flag.String("port", httpserver.DEFAULT_PORT, "server port")
	certFile = flag.String("tlscert", "", "TLS certificate file")
	keyFile  = flag.String("tlskey", "", "TLS private key file")
	verbose  = flag.Bool("verbose", httpserver.DEFAULT_VERBOSE, "verbose output")
	version  = ""
)

func init() {
	log.SetPrefix("dummystore ")
}

func main() {
	flag.Parse()

	adapter := dummyadapter.New(version)
	config := &httpserver.Config{
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
		Verbose:  *verbose,
	}
	h := httpserver.New(adapter, config)

	log.Fatal(h.ListenAndServe())
}
