/*
 * SPDX-License-Identifier: BSD-2-Clause
 *
 * Copyright (c) Lewis Cook <lcook@FreeBSD.org>
 * All rights reserved.
 */
package fetcher

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"sync"

	. "github.com/lcook/portsync/internal/_package"
	. "github.com/lcook/portsync/internal/util"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	baseURL string = "https://portscout.freebsd.org/"
	rssURL         = baseURL + "rss/rss.cgi?m="

	packageDefault string = "DEFAULT"
	packageCargo          = "CARGO"
	packageGo             = "GO"
)

type Portscout struct{}

func (ps Portscout) Fetch(cmd *cobra.Command) (*Packages, error) {
	maintainer := viper.GetString("maintainer")
	data, err := feed(rssURL + maintainer)
	if err != nil {
		return nil, err
	}
	if len(data) < 1 {
		return nil, fmt.Errorf("no package updates found for maintainer '%s'", maintainer)
	}
	var (
		errors   = make(chan error, 1)
		wgroup   sync.WaitGroup
		mutex    sync.Mutex
		packages Packages
	)
	for _, item := range data {
		wgroup.Add(1)
		go func(i *gofeed.Item) {
			defer wgroup.Done()
			entry := feedEntry{i}
			pkg, err := ps.Transform(cmd, &Package{
				Origin: entry.ext("portcat") + "/" + entry.ext("portname"),
				Latest: entry.ext("newversion"),
				Type:   packageDefault,
			})
			if err != nil {
				errors <- err
				return
			}
			mutex.Lock()
			packages.Add(pkg)
			mutex.Unlock()
		}(item)
	}
	go func() {
		wgroup.Wait()
		close(errors)
	}()
	if err := <-errors; err != nil {
		return &Packages{}, err
	}
	packages.Filter(func(pkg Package) bool {
		out, _ := exec.Command("pkg", "version", "-t", pkg.Version, pkg.Latest).Output()
		return pkg.Version == pkg.Latest || strings.TrimRight(string(out), "\n") != "<"
	})
	origins := viper.GetStringSlice("origins")
	if len(origins) > 0 {
		packages.Filter(func(pkg Package) bool {
			return !slices.Contains(origins, pkg.Origin)
		})
	}
	return &packages, nil
}

func (ps Portscout) Transform(cmd *cobra.Command, pkg *Package) (*Package, error) {
	pkgPath := CleanPath(viper.GetString("base")) + pkg.Origin
	if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
		return &Package{}, err
	}
	makeVariable := func(val string) string {
		out, err := exec.Command("make", "-C", pkgPath, "-V", val).Output()
		if err != nil {
			return ""
		}
		return strings.TrimRight(string(out), "\n")
	}
	if strings.Contains(makeVariable("USES"), "cargo") {
		pkg.Type = packageCargo
	}
	if strings.Contains(makeVariable("USES"), "go:modules") &&
		makeVariable("GO_MODULE") == "" {
		pkg.Type = packageGo
	}
	return &Package{
		Origin:     pkg.Origin,
		Version:    makeVariable("DISTVERSION"),
		Latest:     strings.TrimPrefix(pkg.Latest, makeVariable("DISTVERSIONPREFIX")),
		Maintainer: makeVariable("MAINTAINER"),
		Type:       pkg.Type,
	}, nil
}

type feedEntry struct {
	data *gofeed.Item
}

func feed(url string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return feed.Items, nil
}

func (f feedEntry) ext(index string) string { return f.data.Extensions["port"][index][0].Value }
