package http_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	statiks "github.com/janiltonmaciel/statiks/pkg/http"
	check "gopkg.in/check.v1"
)

var noCacheHeaders = map[string]string{
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
	"Expires":         "0",
}

func (s *StatiksSuite) TestNoCacheHandler(c *check.C) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()

	handler := statiks.NoCacheHandler(testHandler)
	handler.ServeHTTP(w, nil)

	for k, v := range noCacheHeaders {
		c.Assert(w.Header().Get(k), check.Equals, v)
	}
}

func (s *StatiksSuite) TestCacheHandler(c *check.C) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	type args struct {
		h     http.Handler
		cache int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Nocache",
			args{testHandler, 0},
			"max-age=0",
		},
		{
			"MaxAge",
			args{testHandler, 99},
			"max-age=99",
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		handler := statiks.CacheHandler(tt.args.h, tt.args.cache)
		handler.ServeHTTP(w, nil)
		c.Assert(w.Header().Get("Cache-Control"), check.Equals, tt.want)
	}
}

func (s *StatiksSuite) TestAddDelayHandler(c *check.C) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()

	handler := statiks.DelayHandler(testHandler, 100*time.Millisecond)
	handler.ServeHTTP(w, nil)

	c.Assert(w.Header().Get("X-Delay"), check.Equals, "100ms")
}
