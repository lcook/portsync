/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"os"
	"os/exec"
	"strings"

	. "github.com/lcook/portsync/internal/fetcher"
	. "github.com/lcook/portsync/internal/util"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: `Run local script or command`,
	Long:  `Run local script or command with inherited package enviroment variables.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packages, err := GetPackages(Portscout{})
		if err != nil {
			return err
		}
		for _, _package := range *packages {
			SetPkgEnv(_package)
			c := exec.Command("sh", "-c", strings.Join(args, " "))
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			_ = c.Run()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
