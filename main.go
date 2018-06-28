package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codegangsta/negroni"
	"github.com/janiltonmaciel/statiks/core"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/cli"

	"github.com/rs/cors"
)

var (
	version string
	commit  string
	date    string
	author  = "Janilton Maciel <janilton@gmail.com>"
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

func init() {
	cli.AppHelpTemplate = core.AppHelpTemplate
	cli.VersionPrinter = core.VersionPrinter(commit, date)
}

func main() {

	app := cli.NewApp()
	app.Name = "statiks"
	app.Usage = "a simple http server to serve static files"
	app.UsageText = "statiks [options]"
	app.Author = author
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, r",
			Value: ".",
			Usage: "The root of the server file tree",
		},

		cli.StringFlag{
			Name:  "port, p",
			Value: "9080",
			Usage: "The port to listen to for incoming HTTP connections",
		},

		cli.BoolFlag{
			Name:  "gzip, z",
			Usage: "Enabled gzip",
		},
	}

	app.Action = action

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func action(c *cli.Context) error {
	path := c.String("path")
	port := c.String("port")
	gZip := c.Bool("gzip")

	docroot, err := filepath.Abs(path)
	if err != nil {
		return err
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
	return nil
}
