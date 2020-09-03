package cmd_test

import (
	"flag"

	"github.com/janiltonmaciel/statiks/cmd"
	"github.com/urfave/cli/v2"
	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestApp(c *check.C) {
	commit := "da3c509"
	date := "2020-09-03T14:45:36Z"
	version := "v0.1"
	app := cmd.CreateApp(version, commit, date)

	c.Assert(app.Version, check.Equals, version)
	c.Assert(app.Authors[0].Name, check.Equals, "Janilton Maciel")
	c.Assert(app.Authors[0].Email, check.Equals, "janilton@gmail.com")
}

func (s *StatiksSuite) TestVersionPrinter(c *check.C) {
	commit := "da3c509"
	date := "2020-09-03T14:45:36Z"
	version := "v0.1"
	app := cmd.CreateApp(version, commit, date)
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)
	cli.VersionPrinter(ctx)
}
