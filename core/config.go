package core

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
	config.host = getHost(c.String("host"))
	config.port = c.String("port")
	config.hidden = c.Bool("hidden")
	config.maxage = c.String("max-age")
	config.origins = getCors(c.String("cors-origins"))
	config.methods = getCors(c.String("cors-methods"))
	config.compress = c.Bool("compress")

	return config
}

var hostReplacer = strings.NewReplacer(
	"http://", "",
	"https://", "",
)

func getHost(host string) string {
	return hostReplacer.Replace(host)
}

func getCors(value string) []string {
	return strings.Split(value, ",")
}
