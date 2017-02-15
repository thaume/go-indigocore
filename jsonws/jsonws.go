// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package jsonws defines functionality to deal with web sockets and JSON.
package jsonws

import "net/http"

// UpgradeHandle is a function that upgrades an HTTP connection to a web socket
// connection.
type UpgradeHandle func(w http.ResponseWriter, r *http.Request, h http.Header) (PingableConn, error)
