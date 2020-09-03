package lib_test

import (
	"flag"
	"github.com/janiltonmaciel/statiks/lib"
	"github.com/urfave/cli/v2"
	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestVersionPrinter(c *check.C) {
	commit := "da3c509"
	date := "2020-09-03T14:45:36Z"
	vp := lib.VersionPrinter(commit, date)

	set := flag.NewFlagSet("test", 0)
	app := cli.NewApp()
	app.Name = "statiks"
	app.Authors = []*cli.Author{{Name: "janilton"}}
	ctx := cli.NewContext(app, set, nil)

	vp(ctx)
}
