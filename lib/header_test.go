package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoCacheHandler(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	w := httptest.NewRecorder()

	handler := NoCacheHandler(testHandler)
	handler.ServeHTTP(w, nil)

	for k, v := range noCacheHeaders {
		assert.Equal(t, w.Header().Get(k), v)
	}
}

func TestCacheHandler(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	type args struct {
		h      http.Handler
		maxAge string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Nocache",
			args{testHandler, "0"},
			"max-age=0",
		},
		{
			"MaxAge",
			args{testHandler, "99"},
			"max-age=99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			handler := CacheHandler(tt.args.h, tt.args.maxAge)
			handler.ServeHTTP(w, nil)
			assert.Equal(t, w.Header().Get("Cache-Control"), tt.want)
		})
	}
}
