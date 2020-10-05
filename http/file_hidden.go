// +build !windows

package lib

import (
	"path/filepath"
	"strings"
)

func hideFile(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}
