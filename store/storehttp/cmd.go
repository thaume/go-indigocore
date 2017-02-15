// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package storehttp

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/Sirupsen/logrus"

	"github.com/stratumn/go/jsonhttp"
	"github.com/stratumn/go/jsonws"
	"github.com/stratumn/go/store"
)

// Run launches a HTTP Store
func Run(
	a store.Adapter,
	httpConfig *jsonhttp.Config,
	basicConfig *jsonws.BasicConfig,
	bufConnConfig *jsonws.BufferedConnConfig,
) {
	log.Info("Copyright (c) 2017 Stratumn SAS")
	log.Info("Apache License 2.0")
	log.Infof("Runtime %s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	h := New(a, httpConfig, basicConfig, bufConnConfig)

	go func() {
		sigc := make(chan os.Signal)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigc
		log.WithField("signal", sig).Info("Got exit signal")
		log.Info("Cleaning up")
		if err := h.Shutdown(); err != nil {
			log.WithField("error", err).Fatal("Failed to shutdown server")
		}
		log.Info("Stopped")
		os.Exit(0)
	}()

	log.WithField("http", httpConfig.Address).Info("Listening")
	if err := h.ListenAndServe(); err != nil {
		log.WithField("error", err).Fatal("Server stopped")
	}
}
