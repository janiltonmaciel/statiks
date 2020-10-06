package statiks_test

import (
	"log"

	"github.com/janiltonmaciel/statiks"
	"github.com/janiltonmaciel/statiks/pkg/http"
)

func ExampleNewServer() {
	config := http.Config{}
	server := statiks.NewServer(config)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
