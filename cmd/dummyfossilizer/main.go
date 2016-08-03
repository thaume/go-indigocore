package main

import (
	"flag"
	"log"

	"github.com/stratumn/go/fossilizer/dummyadapter"
	"github.com/stratumn/go/fossilizer/httpserver"
	"github.com/stratumn/go/jsonhttp"
)

var (
	port             = flag.String("port", httpserver.DEFAULT_PORT, "server port")
	certFile         = flag.String("tlscert", "", "TLS certificate file")
	keyFile          = flag.String("tlskey", "", "TLS private key file")
	numResultWorkers = flag.Int("workers", httpserver.DEFAULT_NUM_RESULT_WORKERS, "number of result workers")
	minDataLen       = flag.Int("mindata", httpserver.DEFAULT_MIN_DATA_LEN, "minimum data length")
	maxDataLen       = flag.Int("maxdata", httpserver.DEFAULT_MAX_DATA_LEN, "maximum data length")
	verbose          = flag.Bool("verbose", httpserver.DEFAULT_VERBOSE, "verbose output")
	version          = ""
)

func init() {
	log.SetPrefix("dummyfossilizer ")
}

func main() {
	flag.Parse()

	a := dummyadapter.New(version)
	c := &httpserver.Config{
		Config: jsonhttp.Config{
			Port:     *port,
			CertFile: *certFile,
			KeyFile:  *keyFile,
			Verbose:  *verbose,
		},
		NumResultWorkers: *numResultWorkers,
		MinDataLen:       *minDataLen,
		MaxDataLen:       *maxDataLen,
	}
	h := httpserver.New(a, c)

	log.Printf("Listening on %s", *port)
	log.Fatal(h.ListenAndServe())
}
