package lib

import (
	"fmt"
	"net/http"
	"time"
)

var noCacheHeaders = map[string]string{
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
	"Expires":         "0",
}

func NoCacheHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func CacheHandler(h http.Handler, maxAge string) http.Handler {
	v := fmt.Sprintf("max-age=%s", maxAge)
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", v)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func DelayHandler(h http.Handler, delay time.Duration) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
