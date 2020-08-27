package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

type Config struct {
	path          string
	host          string
	port          string
	quiet         bool
	delay         time.Duration
	cache         int
	noIndex       bool
	compression   bool
	includeHidden bool
	cors          bool
	ssl           bool
	cert          string
	key           string

	address      string
	hasCache     bool
	firstRequest bool
}

var addressReplacer = strings.NewReplacer(
	"http://", "",
	"https://", "",
)

func NewConfig(c *cli.Context) (config Config) {
	config.host = getHostAddress(c)
	config.path = getPath(c)
	config.port = c.String("port")
	config.quiet = c.Bool("quiet")
	config.delay = getDelay(c)
	config.cache = c.Int("cache")
	config.noIndex = c.Bool("no-index")
	config.compression = c.Bool("compression")
	config.includeHidden = c.Bool("include-hidden")
	config.cors = c.Bool("cors")
	config.ssl = c.Bool("ssl")
	config.cert = c.String("cert")
	config.key = c.String("key")

	config.address = fmt.Sprintf("%s:%s", config.host, config.port)
	config.firstRequest = true
	config.hasCache = false
	if config.cache > 0 {
		config.hasCache = true
	}

	return config
}

func getPath(c *cli.Context) (path string) {
	path = strings.TrimSpace(c.Args().Get(0))
	if path == "" {
		path = "."
	}
	return path
}

func getHostAddress(c *cli.Context) string {
	address := c.String("address")
	return addressReplacer.Replace(address)
}

func getDelay(c *cli.Context) time.Duration {
	delay := c.Int64("delay")
	return time.Duration(delay) * time.Millisecond
}
