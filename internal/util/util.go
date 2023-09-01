package util

import (
	"os"
	"path/filepath"
	"strings"
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
