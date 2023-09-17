//go:build freebsd

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

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run local port(s) test suite",
	Long:  `Attempt to run test suites on port(s) from local ports tree.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packages, err := GetPackages(Portscout{})
		if err != nil {
			return err
		}
		for _, pkg := range *packages {
			mk.Test(pkg)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
