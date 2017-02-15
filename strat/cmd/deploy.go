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

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy <env> [args...]",
	Short: "Deploy project",
	Long: `Deploy project in current directory to an environment.

It executes, if present, the deploy:<env> command of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("expected environment")
		}
		script := fmt.Sprintf(DeployScriptFmt, args[0])
		return runScript(script, "", args[1:], false, useStdin)
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
}
