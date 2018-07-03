# Statiks
[![Build Status](https://travis-ci.org/janiltonmaciel/statiks.svg?branch=master)](https://travis-ci.org/janiltonmaciel/statiks)
[![Go Report Card](https://goreportcard.com/badge/github.com/janiltonmaciel/statiks)](https://goreportcard.com/report/github.com/janiltonmaciel/statiks)
[![GitHub Releases](https://img.shields.io/github/release/janiltonmaciel/statiks.svg)](https://github.com/janiltonmaciel/statiks/releases)

## Installation

#### Binaries

- **darwin (macOS)** [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.1.0/statiks_0.1.0_macOS_amd64.tar.gz)
- **linux** [386](https://github.com/janiltonmaciel/statiks/releases/download/0.1.0/statiks_0.1.0_linux_386.tar.gz) / [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.1.0/statiks_0.1.0_linux_amd64.tar.gz)
- **windows** [386](https://github.com/janiltonmaciel/statiks/releases/download/0.1.0/statiks_0.1.0_windows_386.zip) / [amd64](https://github.com/janiltonmaciel/statiks/releases/download/0.1.0/statiks_0.1.0_windows_amd64.zip)

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
$ statiks [options] path
```

## options
    --host value, -t value          the host (default: "localhost")
    --port value, -p value          the port to listen to for incoming HTTP connections (default: "9080")
    --hidden, -n                    allow transfer of hidden files (default to false)
    --max-age value, -a value       browser cache max-age in milliseconds (default: 0, no-cache)
    --cors-origins value, -o value  a list of origins a cross-domain request can be executed from (default: "*")
    --cors-methods value, -m value  a list of methods the client is allowed to use with cross-domain requests (default: "HEAD, GET, POST, PUT, PATCH, OPTIONS")
    --compress, -c                  enable gzip compression (default to false)
    --help, -h                      show help
    --version, -v                   print the version


## Examples
  - start server at http://localhost:9080 serving "." (current directory)
  ```bash
    $ statiks
  ```

  - start server at http://localhost:9080 serving "~/Projects" with allowed methods "GET, POST"
  ```bash
    $ statiks -m "GET, POST" ~/Projects
  ```

  - start server at http://192.168.1.100:3000 serving "~/Data" with allowed transfer of hidden files
  ```bash
    $ statiks -t 192.168.1.100 -p 3000 --hidden ~/Data
  ```