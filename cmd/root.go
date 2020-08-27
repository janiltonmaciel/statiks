package cmd

import (
	"os"

	"github.com/janiltonmaciel/statiks/lib"
	"github.com/urfave/cli"
)

var author = "Janilton Maciel <janilton@gmail.com>"

func Execute(version, commit, date string) error {
	cli.AppHelpTemplate = lib.AppHelpTemplate
	cli.VersionPrinter = lib.VersionPrinter(commit, date)

	app := createCliApp(
		version,
	)
	return app.Run(os.Args)
}

func createCliApp(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "statiks"
	app.Usage = "fast, zero-configuration, static HTTP filer server."
	app.UsageText = "statiks [options] <path>"
	app.Author = author
	app.Version = version
	app.Action = lib.MainAction
	app.Flags = createFlags()
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "v, version",
		Usage: "print the version",
	}

	return app
}

func createFlags() []cli.Flag {
	flags := []cli.Flag{
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
	return flags
}
