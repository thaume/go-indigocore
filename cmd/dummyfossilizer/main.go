// Copyright 2016 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/stratumn/go/dummyfossilizer"
	"github.com/stratumn/go/fossilizer/fossilizerhttp"
	"github.com/stratumn/go/jsonhttp"
)

var (
	port             = flag.String("port", fossilizerhttp.DefaultPort, "server port")
	certFile         = flag.String("tlscert", "", "TLS certificate file")
	keyFile          = flag.String("tlskey", "", "TLS private key file")
	numResultWorkers = flag.Int("workers", fossilizerhttp.DefaultNumResultWorkers, "number of result workers")
	minDataLen       = flag.Int("mindata", fossilizerhttp.DefaultMinDataLen, "minimum data length")
	maxDataLen       = flag.Int("maxdata", fossilizerhttp.DefaultMaxDataLen, "maximum data length")
	callbackTimeout  = flag.Duration("callbacktimeout", fossilizerhttp.DefaultCallbackTimeout, "callback requests timeout")
	verbose          = flag.Bool("verbose", fossilizerhttp.DefaultVerbose, "verbose output")
	version          = "0.1.0"
	commit           = "00000000000000000000000000000000"
)

func init() {
	log.SetPrefix("dummyfossilizer ")
}

func main() {
	flag.Parse()

	log.Printf("%s v%s@%s", dummyfossilizer.Description, version, commit[:7])
	log.Print("Copyright (c) 2016 Stratumn SAS")
	log.Print("Apache License 2.0")
	log.Printf("Runtime %s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	a := dummyfossilizer.New((&dummyfossilizer.Config{Version: version, Commit: commit}))

	go func() {
		sigc := make(chan os.Signal)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigc
		log.Printf("Got signal %q", sig)
		log.Print("Stopped")
		os.Exit(0)
	}()

	config := &fossilizerhttp.Config{
		NumResultWorkers: *numResultWorkers,
		MinDataLen:       *minDataLen,
		MaxDataLen:       *maxDataLen,
		CallbackTimeout:  *callbackTimeout,
	}
	httpConfig := &jsonhttp.Config{
		Port:     *port,
		CertFile: *certFile,
		KeyFile:  *keyFile,
		Verbose:  *verbose,
	}
	h := fossilizerhttp.New(a, config, httpConfig)

	log.Printf("Listening on %q", *port)
	if err := h.ListenAndServe(); err != nil {
		log.Fatalf("Fatal: %s", err)
	}
}
