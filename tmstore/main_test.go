// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package tmstore

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// start a tendermint node (and tmpop app) in the background to test against
	StartNode()
	os.Exit(m.Run())
}
