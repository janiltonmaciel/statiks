package main

import (
	"fmt"
	golog "log"

	"github.com/janiltonmaciel/statiks/cmd"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	fmt.Printf(">> %s %s %s\n", version, commit, date)
	if err := cmd.Execute(version, commit, date); err != nil {
		golog.Fatal(err)
	}
}
