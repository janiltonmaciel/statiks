package lib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

var projectName = "statiks"

const (
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

func MainAction(c *cli.Context) error {
	config := getStatiksConfig(c)

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

	// add delay
	if config.delay > 0 {
		handler = DelayHandler(handler, config.delay)
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())

	// add middleware logger
	if !config.quiet {
		n.Use(NewLogger(projectName))
	}

	// add middleware gzip
	if config.gzip {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	// enable cors
	if config.cors {
		n.Use(cors.AllowAll())
	}

	n.UseHandler(mux)

	// printStatiksConfig(config)

	if config.ssl {
		return runHTTPS(config, n)
	}

	return runHTTP(config, n)
}

func runHTTP(config statiksConfig, handler http.Handler) error {
	addr := fmt.Sprintf("%s:%s", config.address, config.port)
	fmt.Printf("Running on \n ⚡️ http://%s, serving '%s'\n\n", addr, config.path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return http.ListenAndServe(addr, handler)
}

func runHTTPS(config statiksConfig, handler http.Handler) error {
	addr := fmt.Sprintf("%s:%s", config.address, config.port)
	cert, key := GetMkCert(addr)
	keyPair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		logger.Fatal("Error: Couldn't create key pair")
	}

	var certificates []tls.Certificate
	certificates = append(certificates, keyPair)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		Certificates:             certificates,
	}

	s := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		TLSConfig:    cfg,
	}

	fmt.Printf("Running on \n ⚡️ https://%s, serving '%s'\n\n", addr, config.path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return s.ListenAndServeTLS("", "")
}

// nolint
func printStatiksConfig(config statiksConfig) {
	fmt.Printf("path: %s\n", config.path)
	fmt.Printf("-a, --address: %s\n", config.address)
	fmt.Printf("-p, --port: %s\n", config.port)
	fmt.Printf("-d, --delay: %s\n", config.delay.String())
	fmt.Printf("-s, --ssl: %t\n", config.ssl)
	fmt.Printf("-c, --cache: %s\n", config.maxage)
	fmt.Printf("-q, --quiet: %t\n", config.quiet)
	fmt.Printf("-g, --gzip: %t\n", config.gzip)
	fmt.Printf("--cors: %t\n", config.cors)
	fmt.Printf("--hidden: %t\n", config.hidden)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
}
