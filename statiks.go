package statiks

import (
	http "github.com/janiltonmaciel/statiks/pkg/http"
)

// NewServer return server implements the static http server.
func NewServer(config http.Config) *http.Server {
	return http.NewServer(config)
}
