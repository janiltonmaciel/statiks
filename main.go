package main

import (
	"log"
	"os"

	"github.com/janiltonmaciel/statiks/cmd"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	app := cmd.CreateApp(version, commit, date)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
