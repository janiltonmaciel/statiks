package http

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// Config configure http server command.
type Config struct {
	Address       string
	Path          string
	Host          string
	Port          string
	Quiet         bool
	Delay         time.Duration
	Cache         int
	NoIndex       bool
	Compression   bool
	IncludeHidden bool
	CORS          bool
	SSL           bool
	Cert          string
	Key           string

	HasCache     bool
	firstRequest bool
}

var hostReplacer = strings.NewReplacer(
	"http://", "",
	"https://", "",
)

// NewConfig create config http server command.
func NewConfig(c *cli.Context) (config Config) {
	config.Host = getHostAddress(c)
	config.Path = getPath(c)
	config.Port = c.String("port")
	config.Quiet = c.Bool("quiet")
	config.Delay = getDelay(c)
	config.Cache = c.Int("cache")
	config.NoIndex = c.Bool("no-index")
	config.Compression = c.Bool("compression")
	config.IncludeHidden = c.Bool("include-hidden")
	config.CORS = c.Bool("cors")
	config.SSL = c.Bool("ssl")
	config.Cert = c.String("cert")
	config.Key = c.String("key")

	config.Address = fmt.Sprintf("%s:%s", config.Host, config.Port)
	config.firstRequest = true
	config.HasCache = false
	if config.Cache > 0 {
		config.HasCache = true
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
	host := c.String("host")
	return hostReplacer.Replace(host)
}

func getDelay(c *cli.Context) time.Duration {
	delay := c.Int64("add-delay")
	return time.Duration(delay) * time.Millisecond
}
