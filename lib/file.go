package lib

import (
	"net/http"
	"os"
)

type neuteredFileSystem struct {
	fs     http.FileSystem
	hidden bool
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	// not allowed hidden file
	if !nfs.hidden {
		if hideFile(path) {
			return nil, os.ErrNotExist
		}
	}

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
