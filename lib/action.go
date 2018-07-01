package lib

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

func MainAction(c *cli.Context) error {

	config := getStatiksConfig(c)
	cors := getCors(config)

	docroot, err := filepath.Abs(config.path)
	if err != nil {
		return err
	}

	fmt.Println("*******************************")
	fmt.Printf("host: http://%s:%s\n", config.host, config.port)
	fmt.Printf("path: %s\n", config.path)
	fmt.Printf("hidden: %t\n", config.hidden)
	fmt.Printf("max-age: %s\n", config.maxage)
	fmt.Printf("origins: %s\n", config.origins)
	fmt.Printf("methods: %s\n", config.methods)
	fmt.Printf("compress: %t\n", config.compress)
	fmt.Println("*******************************")

	nss := neuteredFileSystem{
		fs:     http.Dir(docroot),
		hidden: config.hidden,
	}
	fs := FileServer(nss, config)

	var handler http.Handler
	if config.maxage == "0" {
		handler = NoCacheHandler(fs)
	} else {
		handler = CacheHandler(fs, config.maxage)
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	n := negroni.Classic()
	if config.compress {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	n.Use(cors)
	n.UseHandler(mux)
	n.Run(config.host + ":" + config.port)
	return nil
}
