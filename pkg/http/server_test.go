package http_test

import (
	"net/http"

	statiks "github.com/janiltonmaciel/statiks/pkg/http"
	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestServerPathDefault(c *check.C) {
	config := statiks.Config{
		Host: "localhost",
		Port: "9080",
	}
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("server_test.go")
}

func (s *StatiksSuite) TestServerEnabledIndex(c *check.C) {
	config := s.newConfig()
	config.NoIndex = false
	e := s.newHTTPTester(config)

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

	resp = e.GET("/pkg/cmd/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("root.go")
}

func (s *StatiksSuite) TestServerDisabledIndex(c *check.C) {
	config := s.newConfig()
	config.NoIndex = true
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusNotFound)

	resp = e.GET("/README.md").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

func (s *StatiksSuite) TestServerEnabledAddDelay(c *check.C) {
	config := s.newConfig()
	config.AddDelay = 100
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("X-Delay").Equal("100ms")
}

func (s *StatiksSuite) TestServerDisabledAddDelay(c *check.C) {
	config := s.newConfig()
	config.AddDelay = 0
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("X-Delay").Empty()
}

func (s *StatiksSuite) TestServerEnabledCache(c *check.C) {
	config := s.newConfig()
	config.Cache = 10
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Cache-Control").Equal("max-age=10")
	resp.Header("Expires").Empty()
}

func (s *StatiksSuite) TestServerDisabledCache(c *check.C) {
	config := s.newConfig()
	config.Cache = 0
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Cache-Control").Equal("no-cache, private, max-age=0")
	resp.Header("Expires").Equal("0")
}

func (s *StatiksSuite) TestServerEnabledCompression(c *check.C) {
	config := s.newConfig()
	config.Compression = true
	e := s.newHTTPTester(config)

	resp := e.GET("/").WithHeader("Accept-Encoding", "gzip").Expect()
	resp.Status(http.StatusOK)

	resp.Header("Content-Encoding").Equal("gzip")
}

func (s *StatiksSuite) TestServerDisabledCompression(c *check.C) {
	config := s.newConfig()
	config.Compression = false
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Content-Encoding").Empty()
}

func (s *StatiksSuite) TestServerDisabledCompression2(c *check.C) {
	config := s.newConfig()
	config.Compression = false
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Header("Content-Encoding").Empty()
}

func (s *StatiksSuite) TestServerEnabledIncludeHidden(c *check.C) {
	config := s.newConfig()
	config.IncludeHidden = true
	e := s.newHTTPTester(config)

	resp := e.GET("/.gitignore").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

func (s *StatiksSuite) TestServerDisabledIncludeHidden(c *check.C) {
	config := s.newConfig()
	config.IncludeHidden = false
	e := s.newHTTPTester(config)

	resp := e.GET("/.gitignore").Expect()
	resp.Status(http.StatusNotFound)
}

func (s *StatiksSuite) TestServerEnabledCORS(c *check.C) {
	config := s.newConfig()
	config.CORS = true
	e := s.newHTTPTester(config)

	resp := e.OPTIONS("/").Expect()
	resp.Status(http.StatusOK)
}

func (s *StatiksSuite) TestServerEnabledSSL(c *check.C) {
	config := s.newConfig()
	config.SSL = true
	e := s.newHTTPTester(config)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

func (s *StatiksSuite) TestServerRun(c *check.C) {
	config := statiks.Config{
		Host: "localhost",
		Port: "invalid",
		SSL:  false,
	}

	server := statiks.NewServer(config)
	err := server.Run()
	c.Assert(err, check.NotNil)

	config.SSL = true
	server = statiks.NewServer(config)
	err = server.Run()
	c.Assert(err, check.NotNil)
}
