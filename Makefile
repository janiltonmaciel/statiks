.SILENT:

COLOR_RESET = \033[0m
COLOR_COMMAND = \033[36m
COLOR_YELLOW = \033[33m
COLOR_GREEN = \033[32m
COLOR_RED = \033[31m

SHELL = /bin/bash
.DEFAULT_GOAL := help

PROJECT := statiks

GITHUB_TOKEN := $(shell git config --get github.token || echo $$GITHUB_TOKEN)

TAG := `git describe --tags`
DATE := `date -u +"%Y-%m-%dT%H:%M:%SZ"`
COMMIT := ""

LDFLAGS := -X main.version=$(TAG) -X main.commit=$(COMMIT) -X main.date=$(DATE)


git-tag:
	@printf "\n"; \
	read -p "Tag ($(TAG)): "; \
	if [ ! "$$REPLY" ]; then \
		printf "\n${COLOR_RED}"; \
		echo "Invalid tag."; \
		exit 1; \
	fi; \
	TAG=$$REPLY; \
	sed -i.bak -r "s/[0-9]+.[0-9]+.[0-9]+/$$TAG/g" README.md && rm README.md.bak 2>/dev/null; \
	sed -i.bak -r "s/[0-9]+.[0-9]+.[0-9]+$$/$$TAG/g" Dockerfile && rm Dockerfile.bak 2>/dev/null; \
	git commit README.md Dockerfile -m "Update README.md and Dockerfile with release $$TAG"; \
	git tag -s $$TAG -m "$$TAG"

## Build project
build:
	echo "Building $(PROJECT)"
	go build -ldflags "$(LDFLAGS)" -o $(PROJECT) main.go

## Release of the project
release: git-tag
	@if [ ! "$(GITHUB_TOKEN)" ]; then \
		echo "github token should be configurated."; \
		exit 1; \
	fi; \
	export GITHUB_TOKEN=$(GITHUB_TOKEN); \
	goreleaser release --rm-dist; \
	goreleaser release -f .goreleaser-docker-brew.yml --rm-dist; \
	echo "Release - OK"


## Setup of the project
setup:
	@go get -u github.com/alecthomas/gometalinter
	@go get -u github.com/golang/dep/...
	@brew install goreleaser/tap/goreleaser
	@make vendor-install
	gometalinter --install --update

## Install dependencies of the project
vendor-install:
	@dep ensure -v

## Visualizing dependencies status of the project
vendor-status:
	@dep status

lint: ## Run all the linters
	gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=dupl \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		--aggregate \
		./...



## Prints this help
help:
	printf "${COLOR_YELLOW}${PROJECT}\n------\n${COLOR_RESET}"
	awk '/^[a-zA-Z\-\_0-9\.%]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "${COLOR_COMMAND}$$ make %s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort
	printf "\n"