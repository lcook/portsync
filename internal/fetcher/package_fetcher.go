/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package fetcher

import (
	. "github.com/lcook/portsync/internal/_package"
	"github.com/spf13/cobra"
)

type PackageFetcher interface {
	Fetch(*cobra.Command) (*Packages, error)
	Transform(*cobra.Command, *Package) (*Package, error)
}

type defaultFetcher = Portscout

func Get(cmd *cobra.Command, p PackageFetcher) (*Packages, error) {
	if p == nil {
		p = defaultFetcher{}
	}
	return p.Fetch(cmd)

}
