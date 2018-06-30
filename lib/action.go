package lib

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

type neuteredFileSystem struct {
	fs     http.FileSystem
	hidden bool
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	// not allowed hidden file
	if !nfs.hidden {
		base := filepath.Base(path)
		if IsHidden(base) {
			return nil, os.ErrNotExist
		}
	}

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func MainAction(c *cli.Context) error {

	config := getStatiksConfig(c)

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

	mux := http.NewServeMux()
	mux.Handle("/", NoCache(fs))

	n := negroni.Classic()
	if config.compress {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}
	n.Use(Cors())
	n.UseHandler(mux)

	n.Run(config.host + ":" + config.port)
	return nil
}

func Cors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}

var noCacheHeaders = map[string]string{
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

func NoCache(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
