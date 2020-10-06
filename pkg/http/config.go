package http

import (
	"fmt"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var hostReplacer = strings.NewReplacer(
	"http://", "",
	"https://", "",
)

// Config configure http server command.
type Config struct {
	Path          string
	Host          string
	Port          string
	AddDelay      int
	Quiet         bool
	Cache         int
	NoIndex       bool
	Compression   bool
	IncludeHidden bool
	CORS          bool
	SSL           bool
	Cert          string
	Key           string

	Delay        time.Duration
	Address      string
	HasCache     bool
	FirstRequest bool
}

// NewConfig create config http server command.
func NewConfig(c *cli.Context) (config Config) {
	config.Path = c.Args().Get(0)
	config.Host = c.String("host")
	config.Port = c.String("port")
	config.Quiet = c.Bool("quiet")
	config.AddDelay = c.Int("add-delay")
	config.Cache = c.Int("cache")
	config.NoIndex = c.Bool("no-index")
	config.Compression = c.Bool("compression")
	config.IncludeHidden = c.Bool("include-hidden")
	config.CORS = c.Bool("cors")
	config.SSL = c.Bool("ssl")
	config.Cert = c.String("cert")
	config.Key = c.String("key")

	return config
}

func (conf *Config) init() {
	if conf.Host == "" {
		conf.Host = "0.0.0.0"
	} else {
		conf.Host = hostReplacer.Replace(conf.Host)
	}

	conf.Path = strings.TrimSpace(conf.Path)
	if conf.Path == "" {
		conf.Path = "."
	}

	if conf.Port == "" {
		conf.Port = "9080"
	}

	conf.Delay = time.Duration(conf.AddDelay) * time.Millisecond
	conf.Address = fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	conf.FirstRequest = true
	conf.HasCache = false
	if conf.Cache > 0 {
		conf.HasCache = true
	}
}
