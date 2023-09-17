package util

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/lcook/portsync/internal/_package"
	"github.com/spf13/viper"
)

func CleanPath(path string) string {
	tmp := path
	if strings.HasPrefix(path, "~/") {
		dir, _ := os.UserHomeDir()
		tmp = filepath.Join(dir, path[2:])
	}
	if !strings.HasSuffix(path, "/") {
		tmp += "/"
	}
	return tmp
}

var (
	PkgEnv = func(pkg *Package) map[string]string {
		return map[string]string{
			"PACKAGE_ORIGIN":     pkg.Origin,
			"PACKAGE_VERSION":    pkg.Version,
			"PACKAGE_LATEST":     pkg.Latest,
			"PACKAGE_MAINTAINER": pkg.Maintainer,
			"PACKAGE_TYPE":       pkg.Type,
			"PACKAGE_DIR":        CleanPath(viper.GetString("base")) + pkg.Origin,
			"PACKAGE_ROOT":       CleanPath(viper.GetString("base")),
		}
	}
	SetPkgEnv = func(pkg *Package) {
		for k, v := range PkgEnv(pkg) {
			os.Setenv(k, v)
		}
	}
)
