// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import "github.com/spf13/cobra"

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up [args...]",
	Short: "Start project services",
	Long: `Start services defined by project in current directory.

It executes, if present, the up command of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runScript(UpScript, "", args, false, useStdin)
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
