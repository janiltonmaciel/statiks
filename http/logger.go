package lib

import (
	"bytes"
	golog "log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/urfave/negroni"
)

const requestMsg = "completed request"

var logg = golog.New(os.Stdout, "", golog.Ldate|golog.Ltime|golog.Lmicroseconds)

type logger struct {
	AppName string
}

func newLogger(appName string) *logger {
	return &logger{AppName: appName}
}

func (l *logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, r)
	elapsed := time.Since(start)

	res := rw.(negroni.ResponseWriter)
	var buf bytes.Buffer

	buf.WriteString("message: ")
	buf.WriteString(requestMsg)
	buf.WriteString(" method: ")
	buf.WriteString(r.Method)
	buf.WriteString(" status: ")
	buf.WriteString(strconv.Itoa(res.Status()))
	buf.WriteString(" uri: ")
	buf.WriteString(r.RequestURI)
	buf.WriteString(" size: ")
	buf.WriteString(strconv.Itoa(res.Size()))
	buf.WriteString(" took: ")
	buf.WriteString(elapsed.String())
	buf.WriteString(" cachecontrol: ")
	buf.WriteString(rw.Header().Get("Cache-Control"))
	logg.Printf("[%s] %s", l.AppName, buf.String())
}
