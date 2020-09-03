// +build !windows

package lib

import (
	"path/filepath"
	"strings"
)

func HideFile(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}
