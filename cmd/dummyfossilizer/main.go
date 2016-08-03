package main

import (
	"flag"
	"log"

	"github.com/stratumn/go/fossilizer/dummyadapter"
	"github.com/stratumn/go/fossilizer/httpserver"
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

	adapter := dummyadapter.New(version)

	config := &httpserver.Config{
		Port:             *port,
		CertFile:         *certFile,
		NumResultWorkers: *numResultWorkers,
		MinDataLen:       *minDataLen,
		MaxDataLen:       *maxDataLen,
		Verbose:          *verbose,
	}

	h := httpserver.New(adapter, config)

	log.Fatal(h.ListenAndServe())
}
