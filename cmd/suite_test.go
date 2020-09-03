package cmd_test

import (
	"testing"

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
