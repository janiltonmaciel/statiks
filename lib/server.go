package lib

import (
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
	docroot, err := filepath.Abs(config.path)
	if err != nil {
		panic(err)
	}

	nss := neuteredFileSystem{
		fs:     http.Dir(docroot),
		hidden: !config.includeHidden,
	}
	fs := FileServer(nss, config)

	var handler http.Handler
	if config.hasCache {
		handler = CacheHandler(fs, config.cache)
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
	if config.compression {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}

	// enable cors
	if config.cors {
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
	if s.config.ssl {
		return s.runHTTPS()
	}
	return s.runHTTP()
}

func (s *Server) runHTTP() error {
	fmt.Printf("Running on HTTP\n ⚡️ http://%s, serving '%s'\n\n", s.config.address, s.config.path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return http.ListenAndServe(s.config.address, s.handler)

}

func (s *Server) runHTTPS() error {
	fmt.Printf("Running on HTTPS\n ⚡️ https://%s, serving '%s'\n\n", s.config.address, s.config.path)
	fmt.Print("CTRL-C to stop the️ server\n")
	return http.ListenAndServeTLS(s.config.address, s.config.cert, s.config.key, s.handler)
}
