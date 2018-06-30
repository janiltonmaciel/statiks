package core

import (
	"strings"

	"github.com/urfave/cli"
)

type statiksConfig struct {
	directory string
	host      string
	port      string
	hidden    bool
	cache     bool
	origins   []string
	methods   []string
	compress  bool
}

func getStatiksConfig(c *cli.Context) (config statiksConfig) {
	config.directory = c.String("directory")
	config.host = getHost(c.String("host"))
	config.port = c.String("port")
	config.hidden = c.Bool("hidden")
	config.cache = c.Bool("cache")
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
