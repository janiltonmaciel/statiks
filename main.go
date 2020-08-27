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

// nolint
func init() {
	cli.AppHelpTemplate = lib.AppHelpTemplate
	cli.VersionPrinter = lib.VersionPrinter(commit, date)
}

// nolint
func main() {
	app := cli.NewApp()
	app.Name = "statiks"
	app.Usage = "fast, zero-configuration, static HTTP filer server."
	app.UsageText = "statiks [options] <path>"
	app.Author = author
	app.Version = version
	app.Action = lib.MainAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "h, host",
			Value: "0.0.0.0",
			Usage: "host address to bind to",
		},

		cli.StringFlag{
			Name:  "p, port",
			Value: "9080",
			Usage: "port number",
		},

		cli.BoolFlag{
			Name:  "q, quiet",
			Usage: "enable quiet mode, don't output each incoming request",
		},

		cli.Int64Flag{
			Name:  "delay",
			Value: 0,
			Usage: "add delay to responses (in milliseconds)",
		},

		cli.IntFlag{
			Name:  "cache",
			Value: 0,
			Usage: "set cache time (in seconds) for cache-control max-age header",
		},

		cli.BoolFlag{
			Name:  "no-index",
			Usage: "disable directory listings",
		},

		cli.BoolFlag{
			Name:  "compression",
			Usage: "enable gzip compression",
		},

		cli.BoolFlag{
			Name:  "include-hidden",
			Usage: "enable hidden files as normal",
		},

		cli.BoolFlag{
			Name:  "cors",
			Usage: "enable CORS allowing all origins with all standard methods with any header and credentials.",
		},

		cli.BoolFlag{
			Name:  "ssl",
			Usage: "enable https",
		},

		cli.StringFlag{
			Name:  "cert",
			Value: "cert.pem",
			Usage: "path to the ssl cert file",
		},

		cli.StringFlag{
			Name:  "key",
			Value: "key.pem",
			Usage: "path to the ssl key file",
		},
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "v, version",
		Usage: "print the version",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
