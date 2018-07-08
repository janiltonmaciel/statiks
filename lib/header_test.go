package lib

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/rs/cors"
	"github.com/stretchr/testify/assert"
)

func Test_getCors(t *testing.T) {

	type args struct {
		config statiksConfig
	}
	tests := []struct {
		name string
		args args
		want *cors.Cors
	}{
		{
			"Default",
			args{statiksConfig{}},
			cors.New(cors.Options{
				AllowedMethods:   []string{"GET", "POST", "HEAD"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}),
		},
		{
			"Methods",
			args{statiksConfig{
				methods: []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
			}},
			cors.New(cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}),
		},
		{
			"Origins",
			args{statiksConfig{
				origins: []string{"*"},
			}},
			cors.New(cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}),
		},
		{
			"Origins",
			args{statiksConfig{
				origins: []string{"http://localhost"},
			}},
			cors.New(cors.Options{
				AllowedOrigins:   []string{"http://localhost"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCors(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCors() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
