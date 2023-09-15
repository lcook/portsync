/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package fetcher

import (
	. "github.com/lcook/portsync/internal/_package"
)

type PackageFetcher interface {
	Fetch() (*Packages, error)
	Transform(*Package) (*Package, error)
}

func GetPackages(p PackageFetcher) (*Packages, error) { return p.Fetch() }
