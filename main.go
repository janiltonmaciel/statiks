package main

import (
	"log"

	"github.com/janiltonmaciel/statiks/cmd"
)

var (
	version = "0.0.0"
	commit  string
	date    string
)

func main() {
	if err := cmd.Run(version, commit, date); err != nil {
		log.Fatal(err)
	}
}
