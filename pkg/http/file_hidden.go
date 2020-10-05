// +build !windows

package http

import (
	"path/filepath"
	"strings"
)

func hideFile(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}
