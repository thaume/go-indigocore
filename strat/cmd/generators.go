// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/stratumn/go/generator/repo"
)

// generatorsCmd represents the generators command
var generatorsCmd = &cobra.Command{
	Use:   "generators",
	Short: "List available generators",
	Long: `List the available generators.
	
It outputs the name, description, author, version, and license of each generator.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("unexpected arguments")
		}

		path := generatorPath()
		repo := repo.New(path, generatorsOwner, generatorsRepo, ghToken)

		list, err := repo.List(generatorsRef)
		if err != nil {
			return err
		}

		tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		if _, err := fmt.Fprintln(tw, "NAME\tDESCRIPTION\tAUTHOR\tVERSION\tLICENSE"); err != nil {
			return err
		}

		for _, desc := range list {
			if _, err := fmt.Fprintf(
				tw,
				"%s\t%s\t%s\t%s\t%s\n",
				desc.Name,
				desc.Description,
				desc.Author,
				desc.Version,
				desc.License,
			); err != nil {
				return err
			}
		}

		return tw.Flush()
	},
}

func init() {
	RootCmd.AddCommand(generatorsCmd)
}
