package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/codegangsta/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
)

var (
	version string
	commit  string
	date    string
	author  = "Janilton Maciel <janilton@gmail.com>"

	printVersion bool
	port         string
	path         string
	gZip         bool
)

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

func initFlags() {
	flag.BoolVar(&printVersion, "version", false, "Output version")
	flag.StringVar(&port, "port", "9080", "The port to listen to for incoming HTTP connections. Defaults to 9080")
	flag.StringVar(&path, "path", ".", "The root of the server file tree")
	flag.BoolVar(&gZip, "gzip", true, "Enabled gzip")
	flag.Parse()
}

func main() {

	initFlags()

	if printVersion {
		printStatiksVersion()
		return
	}

	docroot, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("*******************************")
	fmt.Printf("$ ./statiks -path=%s -listen=%s  -gzip=%t \n\n", path, port, gZip)
	fmt.Printf("Host: http://localhost:%s\n", port)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Gzip: %t\n", gZip)
	fmt.Println("*******************************")

	fs := http.FileServer(http.Dir(docroot))

	mux := http.NewServeMux()
	mux.Handle("/", NoCache(fs))

	n := negroni.Classic()
	if gZip {
		n.Use(gzip.Gzip(gzip.BestSpeed))
	}
	n.Use(Cors())
	n.UseHandler(mux)

	n.Run(":" + port)
}

func printStatiksVersion() {
	fmt.Println("Version:", version)
	fmt.Println("Author:", author)
	fmt.Println("Commit:", commit)
	fmt.Println("Date:", date)
}
