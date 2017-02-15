// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import "github.com/spf13/cobra"

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [args...]",
	Short: "Run project test suite",
	Long: `Run the test suite of project in current directory.

It executes, if present, the test command of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		testErr := runScript(TestScript, "", args, false, useStdin)
		downErr := runScript(DownTestScript, "", nil, true, useStdin)

		if testErr != nil {
			return testErr
		}

		return downErr
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}
