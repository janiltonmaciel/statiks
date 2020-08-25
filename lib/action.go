package lib

import (
	"github.com/urfave/cli"
)

func MainAction(c *cli.Context) error {
	config := NewConfig(c)
	server := NewServer(config)
	return server.Run()
}
