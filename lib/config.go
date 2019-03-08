package lib

import (
	"strings"
	"time"

	"github.com/urfave/cli"
)

type statiksConfig struct {
	path    string
	address string
	port    string
	hidden  bool
	maxage  string
	cors    bool
	ssl     bool
	delay   time.Duration
	quiet   bool
	gzip    bool

	cache        bool
	firstRequest bool
}

var addressReplacer = strings.NewReplacer(
	"http://", "",
	"https://", "",
)

func getStatiksConfig(c *cli.Context) (config statiksConfig) {
	config.path = getPath(c)
	config.address = getContextAddress(c)
	config.port = c.String("port")
	config.ssl = c.Bool("ssl")
	config.delay = getDelay(c)

	config.hidden = !c.Bool("hidden")
	config.maxage = getContextMaxAge(c)
	config.cors = c.Bool("cors")
	config.gzip = c.Bool("gzip")
	config.quiet = c.Bool("quiet")
	config.cache = config.maxage != "0"
	config.firstRequest = true

	return config
}

func getPath(c *cli.Context) (path string) {
	path = strings.TrimSpace(c.Args().Get(0))
	if path == "" {
		path = "."
	}
	return path
}

func getContextAddress(c *cli.Context) string {
	address := c.String("address")
	return addressReplacer.Replace(address)
}

func getContextMaxAge(c *cli.Context) string {
	value := strings.TrimSpace(c.String("max-age"))
	maxge := "0"
	if value != "" {
		maxge = value
	}
	return maxge
}

func getDelay(c *cli.Context) time.Duration {
	delay := c.Int64("delay")
	return time.Duration(delay) * time.Millisecond
}
