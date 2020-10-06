package statiks_test

import (
	"log"

	"github.com/janiltonmaciel/statiks"
	"github.com/janiltonmaciel/statiks/pkg/http"
)

func ExampleNewServer() {
	config := http.Config{
		Host: "localhost",
		Port: "9080",
	}
	server := statiks.NewServer(config)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
