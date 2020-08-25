# Statiks
Fast, zero-configuration, static HTTP filer server.
Like Python's `SimpleHTTPServer` but scalable.

[![Build Status](https://travis-ci.org/janiltonmaciel/statiks.svg?branch=master)](https://travis-ci.org/janiltonmaciel/statiks)
[![Go Report Card](https://goreportcard.com/badge/github.com/janiltonmaciel/statiks)](https://goreportcard.com/report/github.com/janiltonmaciel/statiks)
[![GitHub Releases](https://img.shields.io/github/release/janiltonmaciel/statiks.svg)](https://github.com/janiltonmaciel/statiks/releases)

## Features

* HTTP and HTTPS (TLS/SSL)
* Support directory list
* Cache control and "Last-Modified"
* Delay response
* CORS support
* Content compression (gzip)
* Request logging

## Installation

### Via Binaries
  * **darwin (macOS)** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.8/statiks_0.8_darwin_amd64.tar.gz)
*  **linux** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.8/statiks_0.8_linux_amd64.tar.gz)
  * **windows** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.8/statiks_0.8_windows_amd64.zip)

### Via Homebrew (macOS)

```bash
brew tap janiltonmaciel/homebrew-tap
brew install statiks
```

### Via Docker

The official janilton/statiks image is available on Docker Hub.
```bash
docker container run -p 9080:9080 -v .:/var/www --rm janilton/statiks
```

> Volume dir: /var/www

> Expose port: 9080


## Use
```bash
statiks [options] <path>

OPTIONS:
  -a value, --address value  host address to bind to (default: "0.0.0.0")
  -p value, --port value     port number (default: "9080")
  -q, --quiet                enable quiet mode, don't output each incoming request
  --delay value              add delay to responses (in milliseconds) (default: 0)
  --cache value              set cache time (in seconds) for cache-control max-age header (default: 0)
  --no-index                 disable directory listings
  --compression              enable gzip compression
  --include-hidden           enable hidden files as normal
  --cors                     enable CORS allowing all origins with all standard methods with any header and credentials.
  --ssl                      enable https
  --cert value               path to the ssl cert file (default: "cert.pem")
  --key value                path to the ssl key file (default: "key.pem")
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
statiks --address 192.168.1.100 --compression /tmp
  ```

  - start server at https://0.0.0.0:9080 serving "." with HTTPS

  ```bash
statiks --ssl
```
