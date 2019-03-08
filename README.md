# Statiks
[![Build Status](https://travis-ci.org/janiltonmaciel/statiks.svg?branch=master)](https://travis-ci.org/janiltonmaciel/statiks)
[![Go Report Card](https://goreportcard.com/badge/github.com/janiltonmaciel/statiks)](https://goreportcard.com/report/github.com/janiltonmaciel/statiks)
[![GitHub Releases](https://img.shields.io/github/release/janiltonmaciel/statiks.svg)](https://github.com/janiltonmaciel/statiks/releases)

## Installation

#### Binaries

- **darwin (macOS)** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.4/statiks_0.4_macOS_amd64.tar.gz)
- **linux** [386](https://github.com/janiltonmaciel/statiks/releases/download/0.4/statiks_0.4_linux_386.tar.gz) / [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.4/statiks_0.4_linux_amd64.tar.gz)
- **windows** [386](https://github.com/janiltonmaciel/statiks/releases/download/0.4/statiks_0.4_windows_386.zip) / [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.4/statiks_0.4_windows_amd64.zip)

#### Via Homebrew (macOS)
```bash
$ brew tap janiltonmaciel/homebrew-tap
$ brew install statiks
```

#### Via Go

```bash
$ go get github.com/janiltonmaciel/statiks
```

## Use
```bash
$ statiks [options] <path>

  -a or --address           set address (default: "0.0.0.0")
  -p or --port              set port (default: "9080")
  -S or --ssl               enable https (default: false)
  -H or --hidden            allow transfer of hidden files (default: false)
  -d or --delay             add delay to responses (milliseconds) (default: 0)
  -c                        set cache time (in seconds) for cache-control max-age header (default: 0)
  --co or --cors-origins    a list of origins a cross-domain request can be executed from (default: "*")
  --cm or --cors-methods    a list of methods the client is allowed to use with cross-domain requests (default: "HEAD, GET, POST, PUT, PATCH, OPTIONS")
  --ng or --no-gzip                    disable GZIP Content-Encoding (default: false)
  -q or --quiet                       quiet mode, don't output each incoming request (default: false)
  -h or --help                        show help
  -v or --version                     print the version
```

## Examples
  start server at http://localhost:9000 serving "." with allowed transfer of hidden files
  ```bash
    $ statiks -port 9000 --hidden
  ```

  - start server at http://localhost:9080 serving "/home" with allowed methods "GET, POST"
  ```bash
    $ statiks --cors-methods "GET, POST" /home
  ```

  - start server at http://192.168.1.100 serving "/tmp" with disable gzip compression
  ```bash
    $ statiks --host 192.168.1.100 --no-gzip /tmp
  ```

  - start server at https://localhost:9080 serving "." with HTTPS
  ```bash
    $ statiks --https
  ```