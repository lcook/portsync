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
	"strings"

	. "github.com/lcook/portsync/internal/_package"
	. "github.com/lcook/portsync/internal/util"
	"github.com/spf13/viper"
)

var pkgEnv = func(pkg *Package) map[string]string {
	return map[string]string{
		"PACKAGE_ORIGIN":     pkg.Origin,
		"PACKAGE_VERSION":    pkg.Version,
		"PACKAGE_LATEST":     pkg.Latest,
		"PACKAGE_MAINTAINER": pkg.Maintainer,
		"PACKAGE_TYPE":       pkg.Type,
		"PACKAGE_DIR":        CleanPath(viper.GetString("base")) + pkg.Origin,
	}
}

func run(pkg *Package, makefile string) {
	for k, v := range pkgEnv(pkg) {
		os.Setenv(k, v)
	}
	scriptDir := CleanPath(viper.GetString("scriptdir"))
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
