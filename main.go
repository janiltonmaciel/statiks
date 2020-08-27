package main

import (
	golog "log"

	"github.com/janiltonmaciel/statiks/cmd"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	if err := cmd.Execute(version, commit, date); err != nil {
		golog.Fatal(err)
	}
}
