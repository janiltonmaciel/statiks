package lib_test

import (
	"net/http"
	"os"
	"time"

	statiks "github.com/janiltonmaciel/statiks/http"
	check "gopkg.in/check.v1"
)

type FakeResponse struct {
	headers http.Header
	body    []byte
	status  int
}

func (r *FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func (s *StatiksSuite) TestWriteNotModified(c *check.C) {
	fr := &FakeResponse{
		headers: map[string][]string{
			"Content-Type":   {"json"},
			"Content-Length": {"100"},
			"Etag":           {"Etagval"},
			"Last-Modified":  {"Lastval"},
		},
	}
	statiks.WriteNotModified(fr)
	c.Assert(fr.Header(), check.HasLen, 1)
	c.Assert(fr.Header()["Etag"][0], check.Equals, "Etagval")
}

func (s *StatiksSuite) TestToHTTPError(c *check.C) {
	tests := []struct {
		input       error
		wantMessage string
		wantStatus  int
	}{
		{input: os.ErrNotExist, wantMessage: "404 page not found", wantStatus: http.StatusNotFound},
		{input: os.ErrPermission, wantMessage: "403 Forbidden", wantStatus: http.StatusForbidden},
		{input: os.ErrInvalid, wantMessage: "500 Internal Server Error", wantStatus: http.StatusInternalServerError},
	}

	for _, tc := range tests {
		msg, status := statiks.ToHTTPError(tc.input)
		c.Assert(msg, check.Equals, tc.wantMessage)
		c.Assert(status, check.Equals, tc.wantStatus)
	}
}

func (s *StatiksSuite) TestCheckIfModifiedSince(c *check.C) {
	tests := []struct {
		inputMethod         string
		inputHeaderModified string
		inputModTime        time.Time
		want                int
	}{
		{inputMethod: "GET", inputHeaderModified: "Mon, 02 Jan 2006 15:04:05 GMT", inputModTime: time.Now(), want: 1},
		{inputMethod: "GET", inputHeaderModified: "", inputModTime: time.Now(), want: 0},
		{inputMethod: "GET", inputHeaderModified: "2323-2323-2323", inputModTime: time.Now(), want: 0},
	}

	for _, tc := range tests {
		config := statiks.Config{}
		req := &http.Request{
			Method: tc.inputMethod,
			Header: map[string][]string{"If-Modified-Since": {tc.inputHeaderModified}},
		}
		result := statiks.CheckIfModifiedSince(req, tc.inputModTime, config)
		c.Assert(int(result), check.Equals, tc.want)
	}
}
