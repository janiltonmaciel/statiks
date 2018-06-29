package core

import (
	"runtime"
)

func IsHidden(filename string) bool {
	if runtime.GOOS == "windows" {
		return isHiddenWindows(filename)
	}

	return isHiddenUnix(filename)
}

func isHiddenUnix(filename string) bool {
	// unix/linux file or directory that starts with . is hidden
	return filename[0:1] == "."
}

func isHiddenWindows(path string) bool {
	// p, e := sys.UTF16PtrFromString(path)
	// if e != nil {
	// 	return false, e
	// }
	// attrs, e := syscall.GetFileAttributes(p)
	// if e != nil {
	// 	return false, e
	// }
	// return attrs & syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil

	return false
}
