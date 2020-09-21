package lib_test

import (
	"flag"
	"net/http"

	"github.com/janiltonmaciel/statiks/lib"
	"github.com/urfave/cli/v2"
	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestServerPathDefault(c *check.C) {
	set := flag.NewFlagSet("test", 0)
	set.String("host", "localhost", "")
	set.String("port", "9080", "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("server_test.go")
}

func (s *StatiksSuite) TestServerEnabledIndex(c *check.C) {
	set := s.newFlagSet()
	set.Bool("no-index", false, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("README.md")

	resp = e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("README.md")

	resp = e.GET("index.html").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("README.md")

	resp = e.GET("/index.html").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("README.md")

	resp = e.GET("/cmd/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("root.go")
}

func (s *StatiksSuite) TestServerDisabledIndex(c *check.C) {
	set := s.newFlagSet()
	set.Bool("no-index", true, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusNotFound)

	resp = e.GET("/README.md").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

func (s *StatiksSuite) TestServerEnabledAddDelay(c *check.C) {
	set := s.newFlagSet()
	set.Int64("add-delay", 100, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("X-Delay").Equal("100ms")
}

func (s *StatiksSuite) TestServerDisabledAddDelay(c *check.C) {
	set := s.newFlagSet()
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("X-Delay").Empty()
}

func (s *StatiksSuite) TestServerEnabledCache(c *check.C) {
	set := s.newFlagSet()
	set.Int("cache", 10, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Cache-Control").Equal("max-age=10")
	resp.Header("Expires").Empty()
}

func (s *StatiksSuite) TestServerDisabledCache(c *check.C) {
	set := s.newFlagSet()
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Cache-Control").Equal("no-cache, private, max-age=0")
	resp.Header("Expires").Equal("0")
}

func (s *StatiksSuite) TestServerEnabledCompression(c *check.C) {
	set := s.newFlagSet()
	set.Bool("compression", true, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").WithHeader("Accept-Encoding", "gzip").Expect()
	resp.Status(http.StatusOK)

	resp.Header("Content-Encoding").Equal("gzip")
}

func (s *StatiksSuite) TestServerDisabledCompression(c *check.C) {
	set := s.newFlagSet()
	set.Bool("compression", true, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Content-Encoding").Empty()
}

func (s *StatiksSuite) TestServerDisabledCompression2(c *check.C) {
	set := s.newFlagSet()
	set.Bool("compression", false, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Content-Encoding").Empty()
}

func (s *StatiksSuite) TestServerEnabledIncludeHidden(c *check.C) {
	set := s.newFlagSet()
	set.Bool("include-hidden", true, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/.gitignore").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

func (s *StatiksSuite) TestServerDisabledIncludeHidden(c *check.C) {
	set := s.newFlagSet()
	set.Bool("include-hidden", false, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/.gitignore").Expect()
	resp.Status(http.StatusNotFound)
}

func (s *StatiksSuite) TestServerEnabledCORS(c *check.C) {
	set := s.newFlagSet()
	set.Bool("cors", true, "")
	e := s.newHTTPTester(set)

	resp := e.OPTIONS("/").Expect()
	resp.Status(http.StatusOK)
}

func (s *StatiksSuite) TestServerEnabledSSL(c *check.C) {
	set := s.newFlagSet()
	set.Bool("ssl", true, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

func (s *StatiksSuite) TestServerRun(c *check.C) {
	set := s.newFlagSet()
	ctx := cli.NewContext(nil, set, nil)
	config := lib.NewConfig(ctx)
	config.Address = "localhost:invalid"
	server := lib.NewServer(config)
	err := server.Run()
	c.Assert(err, check.NotNil)

	config.SSL = true
	config.Address = "localhost:invalid"
	server = lib.NewServer(config)
	err = server.Run()
	c.Assert(err, check.NotNil)
}
