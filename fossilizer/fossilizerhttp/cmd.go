// Copyright 2017 Stratumn SAS. All rights reserved.
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

package fossilizerhttp

import (
	"context"
	"flag"
	"net/http"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/stratumn/go-indigocore/fossilizer"
	"github.com/stratumn/go-indigocore/jsonhttp"
	"github.com/stratumn/go-indigocore/jsonws"
	"github.com/stratumn/go-indigocore/monitoring"
)

var (
	fossilizerEventChanSize int
	addr                    string
	wsReadBufSize           int
	wsWriteBufSize          int
	wsWriteChanSize         int
	wsWriteTimeout          time.Duration
	wsPongTimeout           time.Duration
	wsPingInterval          time.Duration
	wsMaxMsgSize            int64
	certFile                string
	keyFile                 string
	minDataLen              int
	maxDataLen              int
	readTimeout             time.Duration
	writeTimeout            time.Duration
	maxHeaderBytes          int
	shutdownTimeout         time.Duration
)

// Run launches a fossilizerhttp server.
func Run(
	ctx context.Context,
	a fossilizer.Adapter,
	config *Config,
	monitoringConfig *monitoring.Config,
	httpConfig *jsonhttp.Config,
	basicConfig *jsonws.BasicConfig,
	bufConnConfig *jsonws.BufferedConnConfig,
	shutdownTimeout time.Duration,
) {
	log.Info("Copyright (c) 2017 Stratumn SAS")
	log.Info("Apache License 2.0")
	log.Infof("Runtime %s %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	h := New(a, config, httpConfig, basicConfig, bufConnConfig)
	h.exposeMetrics(monitoringConfig)

	go func() {
		<-ctx.Done()
		log.Info("Cleaning up")
		if err := h.Shutdown(ctx); err != nil {
			log.WithField("error", err).Fatal("Failed to shutdown server")
		}
		log.Info("Stopped")
	}()

	log.WithField("http", httpConfig.Address).Info("Listening")
	if err := h.ListenAndServe(); err != http.ErrServerClosed {
		log.WithField("error", err).Fatal("Server stopped")
	}
}

// RegisterFlags register the flags used by RunWithFlags.
func RegisterFlags() {
	flag.IntVar(&fossilizerEventChanSize, "event_chan_size", DefaultFossilizerEventChanSize, "Size of the FossilizerEvent channel")
	flag.StringVar(&addr, "http", DefaultAddress, "HTTP address")
	flag.StringVar(&certFile, "tls_cert", "", "TLS certificate file")
	flag.StringVar(&keyFile, "tls_key", "", "TLS private key file")
	flag.IntVar(&minDataLen, "mindata", DefaultMinDataLen, "Minimum data length")
	flag.IntVar(&maxDataLen, "maxdata", DefaultMaxDataLen, "Maximum data length")
	flag.DurationVar(&readTimeout, "read_timeout", jsonhttp.DefaultReadTimeout, "Read timeout")
	flag.DurationVar(&writeTimeout, "write_timeout", jsonhttp.DefaultWriteTimeout, "Write timeout")
	flag.IntVar(&maxHeaderBytes, "max_header_bytes", jsonhttp.DefaultMaxHeaderBytes, "Maximum header bytes")
	flag.DurationVar(&shutdownTimeout, "shutdown_timeout", 10*time.Second, "Shutdown timeout")
	flag.IntVar(&wsReadBufSize, "ws_read_buf_size", jsonws.DefaultWebSocketReadBufferSize, "Web socket read buffer size")
	flag.IntVar(&wsWriteBufSize, "ws_write_buf_size", jsonws.DefaultWebSocketWriteBufferSize, "Web socket write buffer size")
	flag.IntVar(&wsWriteChanSize, "ws_write_chan_size", jsonws.DefaultWebSocketWriteChanSize, "Size of a web socket connection write channel")
	flag.DurationVar(&wsWriteTimeout, "ws_write_timeout", jsonws.DefaultWebSocketWriteTimeout, "Timeout for a web socket write")
	flag.DurationVar(&wsPongTimeout, "ws_pong_timeout", jsonws.DefaultWebSocketPongTimeout, "Timeout for a web socket expected pong")
	flag.DurationVar(&wsPingInterval, "ws_ping_interval", jsonws.DefaultWebSocketPingInterval, "Interval between web socket pings")
	flag.Int64Var(&wsMaxMsgSize, "max_msg_size", jsonws.DefaultWebSocketMaxMsgSize, "Maximum size of a received web socket message")
}

// RunWithFlags should be called after RegisterFlags and flag.Parse to launch
// a fossilizerhttp server configured using flag values.
func RunWithFlags(ctx context.Context, a fossilizer.Adapter) {
	config := &Config{
		MinDataLen:              minDataLen,
		MaxDataLen:              maxDataLen,
		FossilizerEventChanSize: fossilizerEventChanSize,
	}
	monitoringConfig := monitoring.ConfigurationFromFlags()
	httpConfig := &jsonhttp.Config{
		Address:        addr,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
		CertFile:       certFile,
		KeyFile:        keyFile,
	}
	basicConfig := &jsonws.BasicConfig{
		ReadBufferSize:  wsReadBufSize,
		WriteBufferSize: wsWriteBufSize,
	}
	bufConnConfig := &jsonws.BufferedConnConfig{
		Size:         wsWriteChanSize,
		WriteTimeout: wsWriteTimeout,
		PongTimeout:  wsPongTimeout,
		PingInterval: wsPingInterval,
		MaxMsgSize:   wsMaxMsgSize,
	}

	Run(
		ctx,
		a,
		config,
		monitoringConfig,
		httpConfig,
		basicConfig,
		bufConnConfig,
		shutdownTimeout,
	)
}
