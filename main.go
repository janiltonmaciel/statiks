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
	app.Usage = "a simple http server to serve static files"
	app.UsageText = "statiks [OPTIONS] path"
	app.Author = author
	app.Version = version
	app.Action = lib.MainAction

	app.Flags = []cli.Flag{
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
			Usage: "allow transfer of hidden files (default: false)",
		},

		cli.StringFlag{
			Name:  "max-age, a",
			Usage: "browser cache max-age in milliseconds (default: 0)",
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
			Name:  "https, s",
			Usage: "enable https (default: false)",
		},

		cli.StringFlag{
			Name:  "cert, c",
			Value: "cert.pem",
			Usage: "client certificate file (PEM format)",
		},

		cli.StringFlag{
			Name:  "key, k",
			Value: "key.pem",
			Usage: "private key file (PEM format)",
		},

		cli.BoolFlag{
			Name:  "compress, e",
			Usage: "enable gzip compression (default: false)",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
