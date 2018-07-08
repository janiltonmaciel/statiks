package lib

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/janiltonmaciel/middleware"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

func MainAction(c *cli.Context) error {

	config := getStatiksConfig(c)
	cors := getCors(config)
	logger := middleware.NewLogger(
		"INFO",
		"statiks",
	)

	docroot, err := filepath.Abs(config.path)
	if err != nil {
		return err
	}

	fmt.Println("~~~~~~~~~~~~~~~~~ parameters ~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("host: %s\n", config.host)
	fmt.Printf("port: %s\n", config.port)
	fmt.Printf("path: %s\n", config.path)
	fmt.Printf("hidden: %t\n", config.hidden)
	fmt.Printf("max-age: %s\n", config.maxage)
	fmt.Printf("origins: %s\n", config.origins)
	fmt.Printf("methods: %s\n", config.methods)
	fmt.Printf("https: %t\n", config.https)
	fmt.Printf("cert: %s\n", config.cert)
	fmt.Printf("key: %s\n", config.key)
	fmt.Printf("compress: %t\n", config.compress)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")

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

	n := negroni.New(
		negroni.NewRecovery(),
		logger,
	)
	if config.compress {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	n.Use(cors)
	n.UseHandler(mux)

	if config.https {
		if !Check(config.cert, config.key) {
			fmt.Println("[statiks] Not found cert or key pem")
		} else {
			fmt.Printf("[statiks] Running on https://%s\n\n", config.addr)
			return http.ListenAndServeTLS(config.addr, config.cert, config.key, n)
		}
	}

	fmt.Printf("[statiks] Running on http://%s\n\n", config.addr)
	return http.ListenAndServe(config.addr, n)
}
