package cmd

import (
	"os"

	"github.com/janiltonmaciel/statiks/lib"
	"github.com/urfave/cli/v2"
)

var author = "Janilton Maciel <janilton@gmail.com>"

func Execute(version, commit, date string) error {
	cli.AppHelpTemplate = lib.AppHelpTemplate
	cli.VersionPrinter = lib.VersionPrinter(commit, date)
	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the version",
	}

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
	app.Authors = []*cli.Author{{Name: author}}
	app.Version = version
	app.Flags = createFlags()
	app.Action = func(c *cli.Context) error {
		config := lib.NewConfig(c)
		server := lib.NewServer(config)
		return server.Run()
	}

	return app
}

// nolint
func createFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"h"},
			EnvVars: []string{"HOST"},
			Value:   "0.0.0.0",
			Usage:   "host address to bind to",
		},
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			EnvVars: []string{"PORT"},
			Value:   "9080",
			Usage:   "port number",
		},
		&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "enable quiet mode, don't output each incoming request",
		},
		&cli.Int64Flag{
			Name:  "delay",
			Value: 0,
			Usage: "add delay to responses (in milliseconds)",
		},
		&cli.IntFlag{
			Name:  "cache",
			Value: 0,
			Usage: "set cache time (in seconds) for cache-control max-age header",
		},
		&cli.BoolFlag{
			Name:  "no-index",
			Usage: "disable directory listings",
		},
		&cli.BoolFlag{
			Name:  "compression",
			Usage: "enable gzip compression",
		},
		&cli.BoolFlag{
			Name:  "include-hidden",
			Usage: "enable hidden files as normal",
		},
		&cli.BoolFlag{
			Name:  "cors",
			Usage: "enable CORS allowing all origins with all standard methods with any header and credentials.",
		},
		&cli.BoolFlag{
			Name:  "ssl",
			Usage: "enable https",
		},
		&cli.StringFlag{
			Name:  "cert",
			Value: "cert.pem",
			Usage: "path to the ssl cert file",
		},
		&cli.StringFlag{
			Name:  "key",
			Value: "key.pem",
			Usage: "path to the ssl key file",
		},
	}
	return flags
}
