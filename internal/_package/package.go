/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package _package

type Package struct {
	Origin     string
	Version    string
	Latest     string
	Maintainer string
	Type       string
}

type Packages []*Package

func (p *Packages) Filter(fn func(Package) bool) {
	for i := 0; i < len((*p)); i++ {
		if fn(*(*p)[i]) {
			*p = append((*p)[:i], (*p)[i+1:]...)
			i--
		}
	}
}

func (p *Packages) Add(pkg *Package) { *p = append(*p, pkg) }
