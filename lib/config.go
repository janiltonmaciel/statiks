package lib

import (
	"strings"
	"time"

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
	https    bool
	delay    time.Duration
	quiet    bool

	addr  string
	cache bool
}

func getStatiksConfig(c *cli.Context) (config statiksConfig) {
	config.path = getPath(c)
	config.host = getContextHost(c.String("host"))
	config.port = c.String("port")
	config.https = c.Bool("https")
	config.delay = getDelay(c)

	config.hidden = c.Bool("hidden")
	config.maxage = getContextMaxAge(c.String("max-age"))
	config.origins = getContextCors(c.String("cors-origins"))
	config.methods = getContextCors(c.String("cors-methods"))
	config.compress = !c.Bool("no-gzip")
	config.quiet = c.Bool("quiet")
	config.cache = config.maxage != "0"

	config.addr = config.host + ":" + config.port

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

func getPath(c *cli.Context) (path string) {
	path = c.Args().Get(0)
	if path == "" {
		path = "."
	}
	return path
}

func getDelay(c *cli.Context) time.Duration {
	delay := c.Int64("delay")
	return time.Duration(delay) * time.Millisecond
}
