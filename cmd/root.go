package cmd

import (
	"fmt"
	"os"

	statiks "github.com/janiltonmaciel/statiks/http"
	"github.com/urfave/cli/v2"
)

// Run CLI application.
func Run(version, commit, date string) error {
	cli.AppHelpTemplate = appHelpTemplate
	cli.VersionPrinter = versionPrinter(commit, date)
	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the version",
	}

	app := newApp(
		version,
	)
	return app.Run(os.Args)
}

func newApp(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "statiks"
	app.Usage = "fast, zero-configuration, static HTTP filer server."
	app.UsageText = "statiks [options] <path>"
	app.Version = version
	app.Authors = createAuthors()
	app.Flags = createFlags()
	app.Action = func(c *cli.Context) error {
		config := statiks.NewConfig(c)
		server := statiks.NewServer(config)
		return server.Run()
	}

	return app
}

func createAuthors() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "Janilton Maciel",
			Email: "janilton@gmail.com",
		},
	}
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
			Name:  "add-delay",
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

func versionPrinter(commit, date string) func(c *cli.Context) {
	return func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version: %s\n", c.App.Version)
		fmt.Fprintf(c.App.Writer, "commit: %s\n", commit)
		fmt.Fprintf(c.App.Writer, "date: %s\n", date)
		fmt.Fprintf(c.App.Writer, "author: %s <%s>\n", c.App.Authors[0].Name, c.App.Authors[0].Email)
	}
}
