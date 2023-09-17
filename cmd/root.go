/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	maintainer string = "ports@FreeBSD.org"
	portsdir   string = "/usr/ports"
	config     string
	origins    []string
	excludes   []string
	scriptDir  string = "/usr/local/share/portsync/Mk"
	version    string = "dev"
	rootCmd           = &cobra.Command{
		SilenceUsage: true,
		Use:          "portsync",
		Version:      version,
		Long: `Command-line utility tailored for FreeBSD, focused on management of package
updates, version tracking and streamlined commits of updated packages.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&portsdir, "base", "b", portsdir, "Local ports directory")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "Configuration file for portsync")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display help page for provided [command]")
	rootCmd.PersistentFlags().StringVarP(&maintainer, "maintainer", "m", maintainer, "Package maintainer")
	rootCmd.PersistentFlags().StringSliceVarP(&origins, "origins", "o", []string{}, "Inclusive filter by package origin(s)")
	rootCmd.PersistentFlags().StringSliceVarP(&excludes, "excludes", "e", []string{}, "Exclusive filter by package origin(s)")
	rootCmd.PersistentFlags().StringVarP(&scriptDir, "scriptdir", "s", scriptDir, "Location of Makefile scripts directory")
}

func initConfig() {
	viper.BindPFlag("base", rootCmd.PersistentFlags().Lookup("base"))
	viper.BindPFlag("maintainer", rootCmd.PersistentFlags().Lookup("maintainer"))
	viper.BindPFlag("origins", rootCmd.PersistentFlags().Lookup("origins"))
	viper.BindPFlag("excludes", rootCmd.PersistentFlags().Lookup("excludes"))
	viper.BindPFlag("scriptdir", rootCmd.PersistentFlags().Lookup("scriptdir"))
	if config != "" {
		viper.SetConfigName(config)
	} else {
		viper.SetConfigName(".portsync")
		viper.SetConfigType("toml")
		viper.AddConfigPath("$HOME")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read config file: ", err.Error())
	}
}
