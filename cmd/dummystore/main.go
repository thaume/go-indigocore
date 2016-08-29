// Copyright 2016 Stratumn SAS. All rights reserved.
// Use of this source code is governed by an Apache License 2.0
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/stratumn/go/dummystore"
	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/store/storehttp"
)

var (
	port     = flag.String("port", storehttp.DefaultPort, "server port")
	certFile = flag.String("tlscert", "", "TLS certificate file")
	keyFile  = flag.String("tlskey", "", "TLS private key file")
	verbose  = flag.Bool("verbose", storehttp.DefaultVerbose, "verbose output")
	version  = "0.1.0"
	commit   = "00000000000000000000000000000000"
)

func init() {
	log.SetPrefix("dummystore ")
}

func main() {
	flag.Parse()

	log.Printf("%s v%s@%s", dummystore.Description, version, commit[:6])
	log.Print("Copyright (c) 2016 Stratumn SAS")
	log.Print("Apache License 2.0")
	log.Printf("Runtime %s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	a := dummystore.New(&dummystore.Config{Version: version, Commit: commit})

	go func() {
		sigc := make(chan os.Signal)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigc
		log.Printf("Got signal %q", sig)
		log.Print("Stopped")
		os.Exit(0)
	}()

	c := &jsonhttp.Config{
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
		Verbose:  *verbose,
	}
	h := storehttp.New(a, c)

	log.Printf("Listening on %q", *port)
	if err := h.ListenAndServe(); err != nil {
		log.Fatalf("Fatal: %s", err)
	}
}
