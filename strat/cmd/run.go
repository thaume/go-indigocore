// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <command> [args...]",
	Short: "Run project command",
	Long: `Run a command for the project in current directory.

It executes the given command of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("expected command")
		}
		return runScript(args[0], "", args[1:], false, useStdin)
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
