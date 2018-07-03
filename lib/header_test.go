package lib

import (
	"reflect"
	"testing"

	"github.com/rs/cors"
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
