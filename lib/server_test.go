package lib_test

import (
	"net/http"

	check "gopkg.in/check.v1"
)

func (s *StatiksSuite) TestServerEnabledIndex(c *check.C) {
	set := s.newFlagSet()
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().Contains("README.md")
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


//func (s *StatiksSuite) TestServerEnabledCORS(c *check.C) {
//	set := s.newFlagSet()
//	set.Bool("cors", true, "")
//	e := s.newHTTPTester(set)
//
//	//resp := e.GET("/").WithHeader("Accept-Encoding", "gzip").Expect()
//	resp := e.OPTIONS("/").WithHeader("Access-Control-Request-Method", "*").Expect()
//	fmt.Printf("Headers: %v\n", resp.Headers())
//	resp.Status(http.StatusOK)
//	//resp.Body().NotEmpty()
//}


func (s *StatiksSuite) TestServerEnabledSSL(c *check.C) {
	set := s.newFlagSet()
	set.Bool("ssl", true, "")
	e := s.newHTTPTester(set)

	resp := e.GET("/").Expect()
	resp.Status(http.StatusOK)
	resp.Body().NotEmpty()
}

