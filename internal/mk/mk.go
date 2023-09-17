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

func run(pkg *Package, makefile string) {
	SetPkgEnv(pkg)
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
