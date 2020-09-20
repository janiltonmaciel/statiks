# Statiks
Fast, zero-configuration, static HTTP filer server.
Like Python's `SimpleHTTPServer` but scalable.

[![GitHub Releases](https://img.shields.io/github/release/janiltonmaciel/statiks.svg)](https://github.com/janiltonmaciel/statiks/releases)
[![GoDoc](https://godoc.org/github.com/janiltonmaciel/statiks?status.svg)](https://pkg.go.dev/github.com/janiltonmaciel/statiks?tab=overview)
[![Go Report Card](https://goreportcard.com/badge/github.com/janiltonmaciel/statiks)](https://goreportcard.com/report/github.com/janiltonmaciel/statiks)
[![Build Status](https://travis-ci.org/janiltonmaciel/statiks.svg?branch=master)](https://travis-ci.org/janiltonmaciel/statiks)

## Features

* HTTP and HTTPS (TLS/SSL)
* Support directory list
* Cache control and "Last-Modified"
* Delay response
* CORS support
* Content compression (gzip)
* Request logging

## Installation

#### Via Binaries
  - **darwin (macOS)** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.16/statiks_0.16_darwin_amd64.tar.gz)
  - **linux** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.16/statiks_0.16_linux_amd64.tar.gz)
  - **windows** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.16/statiks_0.16_windows_amd64.zip)

#### Via Homebrew (macOS)

```bash
brew tap janiltonmaciel/homebrew-tap
brew install statiks
```

#### Via Docker

The official [janilton/statiks](https://hub.docker.com/r/janilton/statiks) image is available on Docker Hub.
```bash
docker container run -p 9080:9080 -v .:/var/www --rm janilton/statiks
```

> Volume dir: /var/www

> Expose port: 9080


## Use
```bash
statiks [options] <path>

OPTIONS:
  --host value, -h value  host address to bind to (default: "0.0.0.0") [$HOST]
  --port value, -p value  port number (default: "9080") [$PORT]
  --quiet, -q             enable quiet mode, don't output each incoming request (default: false)
  --add-delay value       add delay to responses (in milliseconds) (default: 0)
  --cache value           set cache time (in seconds) for cache-control max-age header (default: 0)
  --no-index              disable directory listings (default: false)
  --compression           enable gzip compression (default: false)
  --include-hidden        enable hidden files as normal (default: false)
  --cors                  enable CORS allowing all origins with all standard methods with any header and credentials. (default: false)
  --ssl                   enable https (default: false)
  --cert value            path to the ssl cert file (default: "cert.pem")
  --key value             path to the ssl key file (default: "key.pem")
  --help                  show help (default: false)
```

> `<path>` defaults to `.` (relative path to the current directory)

## Examples
  - start server at http://0.0.0.0:9000 serving "." current directory
```bash
statiks -port 9000
```

  - start server at http://0.0.0.0:9080 serving "/home" with CORS
```bash
statiks --cors /home
```

  - start server at http://192.168.1.100:9080 serving "/tmp" with gzip compression
```bash
statiks --host 192.168.1.100 --compression /tmp
```

  - start server at https://0.0.0.0:9080 serving "." with HTTPS

```bash
statiks --ssl --cert cert.pem --key key.pem
```

  - start server at http://0.0.0.0:9080 serving "/tmp" with delay response 100ms

```bash
statiks --add-delay 100 /tmp
```

## Credits

* Check - [go-check/check](https://github.com/go-check/check) (testing)
* Cli - [urfave/cli](https://github.com/urfave/cli)
* Cors - [rs/cors](https://github.com/rs/cors)
* Httpexpect - [gavv/httpexpect](https://github.com/gavv/httpexpect) (testing)
* Negroni - [urfave/negroni](https://github.com/urfave/negroni)
