/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"fmt"
	"strings"

	. "github.com/lcook/portsync/internal/_package"
	. "github.com/lcook/portsync/internal/fetcher"
	"github.com/spf13/cobra"
)

var (
	format string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: `Display latest package update versions`,
	Long: `Leverages portscout to collect the latest package update versions
and then presents them in a user-friendly format.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packages, err := GetPackages(Portscout{})
		if err != nil {
			return err
		}
		format, _ := cmd.Flags().GetString("format")
		for _, _package := range *packages {
			fmt.Println(formatPkg(_package, format))
		}
		return nil
	},
}

func formatPkg(pkg *Package, format string) string {
	specifiers := map[string]string{
		"%o": pkg.Origin,
		"%v": pkg.Version,
		"%l": pkg.Latest,
		"%m": pkg.Maintainer,
	}
	for k, v := range specifiers {
		format = strings.Replace(format, k, v, -1)
	}
	return format
}

func init() {
	fetchCmd.Flags().StringVarP(&format, "format", "f", "%o: %v -> %l", "Package update format output specifier.")
	rootCmd.AddCommand(fetchCmd)
}
