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
	app.UsageText = "statiks [options] [path]"
	app.Author = author
	app.Version = version
	app.Action = lib.MainAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "a, address",
			Value: "0.0.0.0",
			Usage: "set address",
		},

		cli.StringFlag{
			Name:  "p, port",
			Value: "9080",
			Usage: "set port",
		},

		cli.Int64Flag{
			Name:  "d, delay",
			Usage: "add delay to responses (in milliseconds)",
		},

		cli.StringFlag{
			Name:  "c, cache",
			Usage: "set cache time (in seconds) for cache-control max-age header (default: 0)",
		},

		cli.BoolFlag{
			Name:  "g, gzip",
			Usage: "enable GZIP Content-Encoding",
		},

		cli.BoolFlag{
			Name:  "s, ssl",
			Usage: "enable https",
		},

		cli.BoolFlag{
			Name:  "q, quiet",
			Usage: "enable quiet mode, don't output each incoming request",
		},

		cli.BoolFlag{
			Name:  "hidden",
			Usage: "enable exclude directory entries whose names begin with a dot (.)",
		},

		cli.BoolFlag{
			Name:  "cors",
			Usage: "enable CORS allowing all origins with all standard methods with any header and credentials.",
		},
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "h, help",
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
