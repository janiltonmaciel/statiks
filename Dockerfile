# Inspired by https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

ARG GO_VERSION=1.15
ARG GOOS=linux
ARG GOARCH=amd64
ARG CGO_ENABLED=0

# Build 'statiks' as a self-contained, statically linked, minimal image
FROM golang:${GO_VERSION}-alpine as builder
ENV BUILD_DIR=/go/src/github.com/janiltonmaciel/statiks
WORKDIR ${BUILD_DIR}

COPY go.* *.go *.md ./
COPY ./cmd ./cmd/
COPY ./pkg ./pkg/

RUN apk update && apk add --no-cache ca-certificates make bash git && update-ca-certificates
ENV CGO_ENABLED=0 GO111MODULE=on
RUN go mod download

RUN CGO_ENABLED=${CGO_ENABLED:-0} GOOS=${GOOS:-linux} GOARCH=${GOARCH:-amd64} go build -ldflags "-w -s" -o statiks cmd/statiks.go
RUN echo '<!doctype html> <h1><a href="https://github.com/janiltonmaciel/statiks">statiks</a> works</h1><img src="./img.svg"/>' > index.html \
    && echo '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 500 500"></svg>' > img.svg \
 && true

# Assemble 'statiks' minimal image using statically linked build
LABEL maintainer="Janilton Maciel <janilton@gmail.com>"
FROM scratch as release

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/janiltonmaciel/statiks/index.html /var/www/
COPY --from=builder /go/src/github.com/janiltonmaciel/statiks/img.svg /var/www/
COPY --from=builder /go/src/github.com/janiltonmaciel/statiks/statiks /usr/bin/

VOLUME ["/var/www"]
EXPOSE 9080

ENTRYPOINT ["statiks"]
CMD ["--host", "0.0.0.0", "--port", "9080", "--cors", "/var/www/"]
