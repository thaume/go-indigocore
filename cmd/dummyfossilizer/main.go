// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/Sirupsen/logrus"

	"github.com/stratumn/sdk/dummyfossilizer"
	"github.com/stratumn/sdk/fossilizer/fossilizerhttp"
	"github.com/stratumn/sdk/jsonhttp"
)

var (
	http             = flag.String("http", fossilizerhttp.DefaultAddress, "HTTP address")
	certFile         = flag.String("tlscert", "", "TLS certificate file")
	keyFile          = flag.String("tlskey", "", "TLS private key file")
	numResultWorkers = flag.Int("workers", fossilizerhttp.DefaultNumResultWorkers, "number of result workers")
	minDataLen       = flag.Int("mindata", fossilizerhttp.DefaultMinDataLen, "minimum data length")
	maxDataLen       = flag.Int("maxdata", fossilizerhttp.DefaultMaxDataLen, "maximum data length")
	callbackTimeout  = flag.Duration("callbacktimeout", fossilizerhttp.DefaultCallbackTimeout, "callback requests timeout")
	version          = "0.1.0"
	commit           = "00000000000000000000000000000000"
)

func main() {
	flag.Parse()

	log.Infof("%s v%s@%s", dummyfossilizer.Description, version, commit[:7])
	log.Info("Copyright (c) 2017 Stratumn SAS")
	log.Info("Mozilla Public License 2.0")
	log.Infof("Runtime %s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	a := dummyfossilizer.New(&dummyfossilizer.Config{Version: version, Commit: commit})

	go func() {
		sigc := make(chan os.Signal)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigc
		log.WithField("signal", sig).Info("Got exit signal")
		log.Info("Stopped")
		os.Exit(0)
	}()

	config := &fossilizerhttp.Config{
		NumResultWorkers: *numResultWorkers,
		MinDataLen:       *minDataLen,
		MaxDataLen:       *maxDataLen,
		CallbackTimeout:  *callbackTimeout,
	}
	httpConfig := &jsonhttp.Config{
		Address:  *http,
		CertFile: *certFile,
		KeyFile:  *keyFile,
	}
	h := fossilizerhttp.New(a, config, httpConfig)

	log.WithField("http", *http).Info("Listening")
	if err := h.ListenAndServe(); err != nil {
		log.WithField("error", err).Fatal("Server stopped")
	}
}
