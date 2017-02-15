// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import "github.com/spf13/cobra"

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull [args...]",
	Short: "Pull project Docker images",
	Long: `Pull Docker images of project in current directory.

It executes, if present, the pull command of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runScript(PullScript, "", args, false, useStdin)
	},
}

func init() {
	RootCmd.AddCommand(pullCmd)
}
