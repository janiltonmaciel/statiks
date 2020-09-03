package lib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var projectName = "statiks"

type Server struct {
	config  Config
	handler http.Handler
}

func NewServer(config Config) *Server {
	docroot, err := filepath.Abs(config.Path)
	if err != nil {
		panic(err)
	}

	nss := neuteredFileSystem{
		fs:     http.Dir(docroot),
		hidden: config.IncludeHidden,
	}
	fs := FileServer(nss, config)

	var handler http.Handler
	if config.HasCache {
		handler = CacheHandler(fs, config.Cache)
	} else {
		handler = NoCacheHandler(fs)
	}

	// add delay
	if config.Delay > 0 {
		handler = DelayHandler(handler, config.Delay)
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	n := negroni.New()
	n.Use(negroni.NewRecovery())

	// add middleware logger
	if !config.Quiet {
		n.Use(NewLogger(projectName))
	}

	// add middleware gzip
	if config.Compression {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	// enable cors
	if config.CORS {
		n.Use(cors.AllowAll())
	}

	n.UseHandler(mux)

	s := &Server{
		config:  config,
		handler: n,
	}

	return s
}

func (s *Server) Run() error {
	if s.config.SSL {
		return s.runHTTPS()
	}
	return s.runHTTP()
}

func (s *Server) GetHandler() http.Handler {
	return s.handler
}

func (s *Server) runHTTP() error {
	printLogo()
	fmt.Printf("\nRunning on HTTP\n ⚡️ http://%s, serving '%s'\n\n", s.config.Address, s.config.Path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return http.ListenAndServe(s.config.Address, s.handler)
}

func (s *Server) runHTTPS() error {
	printLogo()
	fmt.Printf("\nRunning on HTTPS\n ⚡️ https://%s, serving '%s'\n\n", s.config.Address, s.config.Path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return http.ListenAndServeTLS(s.config.Address, s.config.Cert, s.config.Key, s.handler)
}

// nolint
func (s *Server) runHTTPSMemory() error {
	cert, key := GetMkCert(s.config.Host)

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

	srv := &http.Server{
		Addr:    s.config.Address,
		Handler: s.handler,
		TLSConfig: cfg,
	}

	printLogo()
	fmt.Printf("Running on HTTPS\n ⚡️ https://%s, serving '%s'\n\n", s.config.Address, s.config.Path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return srv.ListenAndServeTLS("", "")
}
