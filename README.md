# Statiks
Fast, zero-configuration, static HTTP filer server.
Like Python's `SimpleHTTPServer` but scalable.

[![Build Status](https://travis-ci.org/janiltonmaciel/statiks.svg?branch=master)](https://travis-ci.org/janiltonmaciel/statiks)
[![Go Report Card](https://goreportcard.com/badge/github.com/janiltonmaciel/statiks)](https://goreportcard.com/report/github.com/janiltonmaciel/statiks)
[![GitHub Releases](https://img.shields.io/github/release/janiltonmaciel/statiks.svg)](https://github.com/janiltonmaciel/statiks/releases)

## Features

* HTTP and HTTPS (TLS)
* CORS support
* Support directory list
* Content compression (auto, gzip, deflate, br)
* Request logging
* Cache control and "Last-Modified"
* Delay response

### Installation

#### Binaries

  * **darwin (macOS)** [amd64](https://github.com/janiltonmaciel/statiks/releases/download0.7/statiks0.7_macOS_amd64.tar.gz)

*  **linux** [386](https://github.com/janiltonmaciel/statiks/releases/download0.7/statiks0.7_linux_386.tar.gz) / [amd64](https://github.com/janiltonmaciel/statiks/releases/download0.7/statiks0.7_linux_amd64.tar.gz)

  * **windows** [386](https://github.com/janiltonmaciel/statiks/releases/download0.7/statiks0.7_windows_386.zip) / [amd64](https://github.com/janiltonmaciel/statiks/releases/download0.7/statiks0.7_windows_amd64.zip)

### Via Homebrew (macOS)

```bash
brew tap janiltonmaciel/homebrew-tap
brew install statiks
```

#### Via Go

```bash
go get github.com/janiltonmaciel/statiks
```

## Use
```bash
statiks [options] <path>

  -a value, --address value  set address (default: "0.0.0.0")
  -p value, --port value     set port (default: "9080")
  -d value, --delay value    add delay to responses (in milliseconds) (default: 0)
  -c value, --cache value    set cache time (in seconds) for cache-control max-age header (default: 0)
  -g, --gzip                 enable GZIP Content-Encoding
  -q, --quiet                enable quiet mode, don't output each incoming request
  --hidden                   enable exclude directory entries whose names begin with a dot (.)
  --cors                     enable CORS allowing all origins with all standard methods with any header and credentials.
  -h, --help                 show help
  -v, --version              print the version
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
statiks --host 192.168.1.100 --gzip /tmp
  ```


  <!-- > Install [`mkcert`](https://github.com/FiloSottile/mkcert#installation) and run `mkcert -install`
  - start server at https://0.0.0.0:9080 serving "." with HTTPS

  ```bash
    $ statiks --ssl
  ``` -->
