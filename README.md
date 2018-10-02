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

  --host value, -H value            set host (default: "localhost")
  --port value, -p value            set port (default: "9080")
  --https                           enable https (default: false)
  --hidden                          allow transfer of hidden files (default: false)
  --delay value, -d value           add delay to responses (ms) (default: 0)
  --max-age value, --ma value       browser cache control max-age in milliseconds (default: 0)
  --cors-origins value, --co value  a list of origins a cross-domain request can be executed from (default: "*")
  --cors-methods value, --cm value  a list of methods the client is allowed to use with cross-domain requests (default: "HEAD, GET, POST, PUT, PATCH, OPTIONS")
  --no-gzip, --ng                   disable GZIP Content-Encoding (default: false)
  --quiet, -q                       quiet mode, don't output each incoming request (default: false)
  --help, -h                        show help
  --version, -v                     print the version
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