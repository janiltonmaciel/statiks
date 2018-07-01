package lib

import (
	"strings"

	"github.com/urfave/cli"
)

type statiksConfig struct {
	path     string
	host     string
	port     string
	hidden   bool
	maxage   string
	origins  []string
	methods  []string
	compress bool
}

func getStatiksConfig(c *cli.Context) (config statiksConfig) {
	if c.Args().Get(0) == "" {
		config.path = "."
	} else {
		config.path = c.Args().Get(0)
	}
	config.host = getContextHost(c.String("host"))
	config.port = c.String("port")
	config.hidden = c.Bool("hidden")
	config.maxage = getContextMaxAge(c.String("max-age"))
	config.origins = getContextCors(c.String("cors-origins"))
	config.methods = getContextCors(c.String("cors-methods"))
	config.compress = c.Bool("compress")

	return config
}

var hostReplacer = strings.NewReplacer(
	"http://", "",
	"https://", "",
)

func getContextHost(host string) string {
	return hostReplacer.Replace(host)
}

func getContextCors(value string) []string {
	return strings.Split(value, ",")
}

func getContextMaxAge(value string) string {
	maxge := "0"
	if value != "" {
		maxge = value
	}
	return maxge
}
