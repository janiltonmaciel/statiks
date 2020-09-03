// +build windows

package lib

import (
	"syscall"
	"path/filepath"
)

func HideFile(path string) bool {
	base := filepath.Base(path)
	filenameW, err := syscall.UTF16PtrFromString(base)
	if err != nil {
		return false
	}

	attrs, e := syscall.GetFileAttributes(filenameW)
	if e != nil {
		return false
	}
	return attrs & syscall.FILE_ATTRIBUTE_HIDDEN != 0
}