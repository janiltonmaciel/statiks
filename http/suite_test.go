package lib_test

import (
	"flag"
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	statiks "github.com/janiltonmaciel/statiks/http"
	"github.com/urfave/cli/v2"
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

func (s *StatiksSuite) newHTTPTester(set *flag.FlagSet) *httpexpect.Expect {
	ctx := cli.NewContext(nil, set, nil)
	config := statiks.NewConfig(ctx)
	server := statiks.NewServer(config)
	baseURL := fmt.Sprintf("http://%s", config.Address)
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
func (s *StatiksSuite) newFlagSet(paths ...string) *flag.FlagSet {
	set := flag.NewFlagSet("test", 0)
	set.String("host", "localhost", "")
	set.String("port", "9080", "")

	if len(paths) == 0 {
		paths = []string{".."}
	}
	if err := set.Parse(paths); err != nil {
		s.t.Error(err)
	}
	return set
}
