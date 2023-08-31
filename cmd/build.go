/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	. "github.com/lcook/portsync/internal/fetcher"
	"github.com/lcook/portsync/internal/mk"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build local port(s)",
	Long:  `Attempt to build port(s) from local ports tree.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packages, err := Get(cmd, nil)
		if err != nil {
			return err
		}
		for _, pkg := range *packages {
			mk.Build(pkg)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
