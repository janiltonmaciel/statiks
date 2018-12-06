# Inspired by https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

ARG GOOS=linux
ARG GOARCH=amd64
ARG CGO_ENABLED=0

# Build 'statiks' as a self-contained, statically linked, minimal image
FROM golang:1.11 as builder
RUN go get -u github.com/golang/dep/...
ENV BUILD_DIR=/go/src/github.com/halverneus/static-file-server
RUN mkdir -p ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
COPY Gopkg.* *.go *.md ./
COPY ./lib ./lib/

RUN dep ensure -v

COPY Makefile ./
RUN CGO_ENABLED=${CGO_ENABLED:-0} GOOS=${GOOS:-linux} GOARCH=${GOARCH:-amd64} make build_optimized
RUN mkdir -p /out /out/www \
 && cp ./statiks /out/statiks \
 && echo '<!doctype html> <h1><a href="https://github.com/janiltonmaciel/statiks">statiks</a> works</h1><img src="./img.svg"/>' > /out/www/index.html \
 && echo '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 500 500"></svg>' > /out/www/img.svg \
 && true


# Assemble 'statiks' minimal image using statically linked build
FROM scratch as release
COPY --from=builder /out/statiks /
ENTRYPOINT ["/statiks"]
USER 1000
## Serve all requests with a max-age of 15 minutes
CMD ["--host", "0.0.0.0", "--max-age", "900000", "/www"]

FROM release as demo
COPY --from=builder /out/www /www/
