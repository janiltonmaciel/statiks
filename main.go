package main

import (
	"log"
	"os"

	"github.com/janiltonmaciel/statiks/lib"
	"github.com/urfave/cli"
)

var (
	version string
	commit  string
	date    string
	author  = "Janilton Maciel <janilton@gmail.com>"
)

func init() {
	cli.AppHelpTemplate = lib.AppHelpTemplate
	cli.VersionPrinter = lib.VersionPrinter(commit, date)
}

func main() {

	app := cli.NewApp()
	app.Name = "statiks"
	app.Usage = "a simple http server"
	app.UsageText = "statiks [OPTIONS] <path>"
	app.Author = author
	app.Version = version
	app.Action = lib.MainAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host, H",
			Value: "localhost",
			Usage: "set host",
		},

		cli.StringFlag{
			Name:  "port, p",
			Value: "9080",
			Usage: "set port",
		},

		cli.BoolFlag{
			Name:  "https",
			Usage: "enable https (default: false)",
		},

		cli.BoolFlag{
			Name:  "hidden",
			Usage: "allow transfer of hidden files (default: false)",
		},

		cli.Int64Flag{
			Name:  "delay, d",
			Usage: "add delay to responses (ms)",
		},

		cli.StringFlag{
			Name:  "max-age, ma",
			Usage: "browser cache control max-age in milliseconds (default: 0)",
		},

		cli.StringFlag{
			Name:  "cors-origins, co",
			Value: "*",
			Usage: "a list of origins a cross-domain request can be executed from",
		},

		cli.StringFlag{
			Name:  "cors-methods, cm",
			Value: "HEAD, GET, POST, PUT, PATCH, OPTIONS",
			Usage: "a list of methods the client is allowed to use with cross-domain requests",
		},

		cli.BoolFlag{
			Name:  "no-gzip, ng",
			Usage: "disable GZIP Content-Encoding (default: false)",
		},

		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "quiet mode, don't output each incoming request (default: false)",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
