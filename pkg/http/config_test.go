package http_test

import (
	"flag"

	statiks "github.com/janiltonmaciel/statiks/pkg/http"
	"github.com/urfave/cli/v2"
	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestConfig(c *check.C) {
	set := flag.NewFlagSet("test", 0)
	set.String("host", "localhost", "")
	set.String("port", "1080", "")
	set.Bool("quiet", true, "")
	set.Int64("add-delay", 100, "")
	set.Int("cache", 10, "")
	set.Bool("no-index", true, "")
	set.Bool("compression", true, "")
	set.Bool("include-hidden", true, "")
	set.Bool("cors", true, "")
	set.Bool("ssl", true, "")
	set.String("cert", "cert123.pem", "")
	set.String("key", "key123.pem", "")
	err := set.Parse([]string{"path"})
	c.Assert(err, check.IsNil)

	ctx := cli.NewContext(nil, set, nil)
	config := statiks.NewConfig(ctx)

	c.Assert(config.Host, check.Equals, "localhost")
	c.Assert(config.Port, check.Equals, "1080")
	c.Assert(config.Quiet, check.Equals, true)
	c.Assert(config.AddDelay, check.Equals, 100)
	c.Assert(config.Cache, check.Equals, 10)
	c.Assert(config.NoIndex, check.Equals, true)
	c.Assert(config.Compression, check.Equals, true)
	c.Assert(config.IncludeHidden, check.Equals, true)
	c.Assert(config.CORS, check.Equals, true)
	c.Assert(config.SSL, check.Equals, true)
	c.Assert(config.Cert, check.Equals, "cert123.pem")
	c.Assert(config.Key, check.Equals, "key123.pem")
	c.Assert(config.Path, check.Equals, "path")
}
