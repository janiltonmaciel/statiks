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
	app.UsageText = "statiks [OPTIONS]"
	app.Author = author
	app.Version = version
	app.Action = core.MainAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "directory, d",
			Value: ".",
			Usage: "the directory of static file",
		},

		cli.StringFlag{
			Name:  "host, t",
			Value: "localhost",
			Usage: "the host",
		},

		cli.StringFlag{
			Name:  "port, p",
			Value: "9080",
			Usage: "the port to listen to for incoming HTTP connections",
		},

		cli.BoolFlag{
			Name:  "hidden, n",
			Usage: "allow transfer of hidden files (default to false)",
		},

		cli.BoolFlag{
			Name:  "cache, z",
			Usage: "enable cache (http 304) responses  (default to false)",
		},

		cli.StringFlag{
			Name:  "cors-origins, o",
			Value: "*",
			Usage: "a list of origins a cross-domain request can be executed from",
		},

		cli.StringFlag{
			Name:  "cors-methods, m",
			Value: "HEAD, GET, POST, PUT, PATCH, OPTIONS",
			Usage: "a list of methods the client is allowed to use with cross-domain requests",
		},

		cli.BoolFlag{
			Name:  "compress, c",
			Usage: "enable gzip compression (default to false)",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
