package main

import (
	"log"

	"github.com/janiltonmaciel/statiks/cmd"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	if err := cmd.Run(version, commit, date); err != nil {
		log.Fatal(err)
	}
}
