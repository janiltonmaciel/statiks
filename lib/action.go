package lib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/janiltonmaciel/middleware"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

var (
	logLevel    = "INFO"
	projectName = "statiks"
	logger      *logrus.Logger
)

func init() {
	logger = logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.999",
	}
}

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

	// add delay
	if config.delay > 0 {
		handler = DelayHandler(handler, config.delay)
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())

	if !config.quiet {
		// add middleware logger
		n.Use(middlewareLogger)
	}

	if config.compress {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	n.Use(cors)
	n.UseHandler(mux)

	printStatiksConfig(config)

	if config.https {
		return runHTTPS(config, n)
	}

	logger.Printf("Running on http://%s ⚡️", config.addr)
	return http.ListenAndServe(config.addr, n)
}

func runHTTPS(config statiksConfig, n *negroni.Negroni) error {
	cert, key := GetMkCert(config.host)
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
		Addr:           config.addr,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      cfg,
	}

	logger.Printf("Running on https://%s ⚡️", config.addr)
	return s.ListenAndServeTLS("", "")
}

func printStatiksConfig(config statiksConfig) {
	fmt.Println("~~~~~~~~~~~~~~~~~ parameters ~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("https: %t\n", config.https)
	fmt.Printf("host: %s\n", config.host)
	fmt.Printf("port: %s\n", config.port)
	fmt.Printf("path: %s\n", config.path)
	fmt.Printf("delay: %s\n", config.delay.String())
	fmt.Printf("hidden: %t\n", config.hidden)
	fmt.Printf("max-age: %s\n", config.maxage)
	fmt.Printf("origins: %s\n", config.origins)
	fmt.Printf("methods: %s\n", config.methods)
	fmt.Printf("quiet: %t\n", config.quiet)
	fmt.Printf("compress: %t\n", config.compress)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
}
