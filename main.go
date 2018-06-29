package main

import (
	"log"
	"os"

	"github.com/janiltonmaciel/statiks/core"
	"github.com/urfave/cli"
)

var (
	version string
	commit  string
	date    string
	author  = "Janilton Maciel <janilton@gmail.com>"
)

func init() {
	cli.AppHelpTemplate = core.AppHelpTemplate
	cli.VersionPrinter = core.VersionPrinter(commit, date)
}

func main() {

	app := cli.NewApp()
	app.Name = "statiks"
	app.Usage = "a simple http server to serve static files"
	app.UsageText = "statiks [options]"
	app.Author = author
	app.Version = version
	app.Action = core.MainAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, t",
			Value: ".",
			Usage: "the root of the server file tree",
		},

		cli.StringFlag{
			Name:  "port, p",
			Value: "9080",
			Usage: "the port to listen to for incoming HTTP connections",
		},

		cli.BoolFlag{
			Name:  "gzip, z",
			Usage: "enable gzip compression (default to false)",
		},

		cli.BoolFlag{
			Name: "hidden,	",
			Usage: "allow transfer of hidden files (default to false)",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
