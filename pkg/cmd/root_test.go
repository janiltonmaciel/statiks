package cmd_test

import (
	"flag"

	"github.com/janiltonmaciel/statiks/pkg/cmd"
	"github.com/urfave/cli/v2"
	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestApp(c *check.C) {
	version := "v0.1"
	app := cmd.NewApp(version)

	c.Assert(app.Version, check.Equals, version)
	c.Assert(app.Authors[0].Name, check.Equals, "Janilton Maciel")
	c.Assert(app.Authors[0].Email, check.Equals, "janilton@gmail.com")
}

func (s *StatiksSuite) TestVersionPrinter(c *check.C) {
	version := "v0.2"
	app := cmd.NewApp(version)

	commit := "da3c509"
	date := "2020-09-03T14:45:36Z"
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)
	vp := cmd.VersionPrinter(commit, date)
	vp(ctx)
}

func (s *StatiksSuite) TestRun(c *check.C) {
	version := "v0.3"
	commit := "da3c509"
	date := "2020-09-03T14:45:36Z"
	err := cmd.Run(version, commit, date)
	c.Assert(err, check.NotNil)
}

func (s *StatiksSuite) TestAction(c *check.C) {
	version := "v0.4"
	app := cmd.NewApp(version)

	set := flag.NewFlagSet("test", 0)
	set.String("host", "localhost2332", "")

	ctx := cli.NewContext(app, set, nil)

	err := app.Action(ctx)
	c.Assert(err, check.NotNil)
}
