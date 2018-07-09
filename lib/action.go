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

var logLevel = "INFO"
var projectName = "statiks"

func MainAction(c *cli.Context) error {

	config := getStatiksConfig(c)
	cors := getCors(config)
	middlewareLogger := middleware.NewLogger(
		logLevel,
		projectName,
	)

	docroot, err := filepath.Abs(config.path)
	if err != nil {
		return err
	}

	nss := neuteredFileSystem{
		fs:     http.Dir(docroot),
		hidden: config.hidden,
	}
	fs := FileServer(nss, config)

	var handler http.Handler
	if config.cache {
		handler = CacheHandler(fs, config.maxage)
	} else {
		handler = NoCacheHandler(fs)
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())

	if !config.quiet {
		n.Use(middlewareLogger)
	}

	if config.compress {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	n.Use(cors)
	n.UseHandler(mux)

	printStatiksConfig(config)

	if config.https {
		if !Check(config.cert, config.certKey) {
			fmt.Printf("[%s] Not found cert or key pem ⚠️\n", projectName)
		} else {
			fmt.Printf("[%s] Running on https://%s ⚡️\n\n", projectName, config.addr)
			return http.ListenAndServeTLS(config.addr, config.cert, config.certKey, n)
		}
	}

	fmt.Printf("[%s] ️Running on http://%s ⚡️\n\n", projectName, config.addr)
	return http.ListenAndServe(config.addr, n)
}

func printStatiksConfig(config statiksConfig) {
	fmt.Println("~~~~~~~~~~~~~~~~~ parameters ~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("https: %t\n", config.https)
	fmt.Printf("host: %s\n", config.host)
	fmt.Printf("port: %s\n", config.port)
	fmt.Printf("path: %s\n", config.path)
	fmt.Printf("hidden: %t\n", config.hidden)
	fmt.Printf("max-age: %s\n", config.maxage)
	fmt.Printf("origins: %s\n", config.origins)
	fmt.Printf("methods: %s\n", config.methods)
	if config.https {
		fmt.Printf("cert: %s\n", config.cert)
		fmt.Printf("cert-key: %s\n", config.certKey)
	}
	fmt.Printf("quiet: %t\n", config.quiet)
	fmt.Printf("compress: %t\n", config.compress)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
}
