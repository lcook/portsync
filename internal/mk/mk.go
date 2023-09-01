/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package mk

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/lcook/portsync/internal/_package"
	"github.com/spf13/viper"
)

func run(pkg *Package, makefile string) {
	base := viper.GetString("base")
	if strings.HasPrefix(base, "~/") {
		dir, _ := os.UserHomeDir()
		base = filepath.Join(dir, base[2:])
	}
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}
	for k, v := range map[string]string{
		"PACKAGE_ORIGIN":     pkg.Origin,
		"PACKAGE_VERSION":    pkg.Version,
		"PACKAGE_LATEST":     pkg.Latest,
		"PACKAGE_MAINTAINER": pkg.Maintainer,
		"PACKAGE_TYPE":       pkg.Type,
		"PACKAGE_DIR":        base + pkg.Origin,
	} {
		os.Setenv(k, v)
	}
	scriptDir := viper.GetString("scriptdir")
	if strings.HasPrefix(scriptDir, "~/") {
		dir, _ := os.UserHomeDir()
		base = filepath.Join(dir, scriptDir[2:])
	}
	if !strings.HasSuffix(scriptDir, "/") {
		scriptDir += "/"
	}
	args := strings.Fields("make -f" + scriptDir + makefile)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func Build(pkg *Package)  { run(pkg, "build.mk") }
func Commit(pkg *Package) { run(pkg, "commit.mk") }
func Test(pkg *Package)   { run(pkg, "test.mk") }
func Update(pkg *Package) { run(pkg, "update.mk") }
