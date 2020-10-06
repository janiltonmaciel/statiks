package http_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	statiks "github.com/janiltonmaciel/statiks/pkg/http"
	check "gopkg.in/check.v1"
)

type StatiksSuite struct {
	t *testing.T
}

var sSuite = &StatiksSuite{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	sSuite.t = t
	check.TestingT(t)
}

var _ = check.Suite(sSuite)

func (s *StatiksSuite) newHTTPTester(config statiks.Config) *httpexpect.Expect {
	server := statiks.NewServer(config)
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	baseURL := fmt.Sprintf("http://%s", address)
	handler := server.GetHandler()
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL: baseURL,
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(s.t),
	})
	return e
}

// nolint
func (s *StatiksSuite) newConfig(paths ...string) statiks.Config {
	config := statiks.Config{
		Host: "localhost",
		Port: "9080",
		Path: "../..",
	}
	return config
}
