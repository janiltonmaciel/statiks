package statiks_test

import (
	"log"

	"github.com/janiltonmaciel/statiks"
	"github.com/janiltonmaciel/statiks/pkg/http"
)

func ExampleNewServer() {
	config := http.Config{}
	server, err := statiks.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
