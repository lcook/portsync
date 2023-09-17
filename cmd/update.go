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

var (
	commit bool = false
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Apply latest package version updates to local ports tree",
	Long: `Fetches the most recent package version information and applies
any necessary updates to your local ports tree, ensuring that it
remains in sync with the latest upstream as determined by portscout.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packages, err := GetPackages(Portscout{})
		if err != nil {
			return err
		}
		for _, pkg := range *packages {
			mk.Update(pkg)
			if commit {
				mk.Commit(pkg)
			}
		}
		return nil
	},
}

func init() {
	updateCmd.Flags().BoolVarP(&commit, "commit", "g", commit, "Commit package update.")
	rootCmd.AddCommand(updateCmd)
}
