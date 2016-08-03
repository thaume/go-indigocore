// A store HTTP server with a dummy adapter.
package main

import (
	"flag"
	"log"

	"github.com/stratumn/go/dummystore"
	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/storehttp"
)

var (
	port     = flag.String("port", storehttp.DefaultPort, "server port")
	certFile = flag.String("tlscert", "", "TLS certificate file")
	keyFile  = flag.String("tlskey", "", "TLS private key file")
	verbose  = flag.Bool("verbose", storehttp.DefaultVerbose, "verbose output")
	version  = ""
)

func init() {
	log.SetPrefix("dummystore ")
}

func main() {
	flag.Parse()

	a := dummystore.New(version)
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
