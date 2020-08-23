package lib

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type neuteredFileSystem struct {
	fs     http.FileSystem
	hidden bool
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	// not allowed hidden file
	if !nfs.hidden {
		base := filepath.Base(path)
		if IsHidden(base) {
			return nil, os.ErrNotExist
		}
	}

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

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

func Check(certPath string, keyPath string) bool {
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return false
	} else if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return false
	}
	return true
}
