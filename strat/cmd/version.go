// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version info",
	Long: `Display the version info of Stratumn CLI.

It outputs the semver string, and the first seven characters of the Git hash.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("unexpected arguments")
		}
		fmt.Printf("%s v%s@%s\n", "Stratumn CLI", version, commit[:7])
		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
