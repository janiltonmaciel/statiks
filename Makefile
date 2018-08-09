.SILENT:

SHELL = /bin/bash
.DEFAULT_GOAL := help

COLOR_RESET = \033[0m
COLOR_COMMAND = \033[36m
COLOR_YELLOW = \033[33m
COLOR_GREEN = \033[32m
COLOR_RED = \033[31m

TAG := `git describe --tags`
DATE := `date -u +"%Y-%m-%dT%H:%M:%SZ"`
COMMIT := ""
LDFLAGS := -X main.version=$(TAG) -X main.commit=$(COMMIT) -X main.date=$(DATE)
GITHUB_TOKEN := $(shell git config --get github.token || echo $$GITHUB_TOKEN)

PROJECT := statiks
SOURCE_FILES?=$$(go list ./... | grep -v /vendor/)

## Run project http
run:
	@go run main.go

## Run project https
runs:
	@go run main.go -s

## Runs the project unit tests
test:
	@go test -timeout 10s  -v -covermode atomic -cover -coverprofile coverage.txt $(SOURCE_FILES)
	@go tool vet . 2>&1 | grep -v '^vendor\/' | grep -v '^exit\ status\ 1' || true

## Run all the tests and opens the coverage report
test-cover: test 
	go tool cover -html=coverage.txt
	@rm coverage.txt 2>/dev/null || true

## Run all the tests and code checks
test-ci: lint test 

## Setup of the project
setup:
	@go get -u github.com/alecthomas/gometalinter
	@go get -u github.com/golang/dep/...
	@brew install goreleaser/tap/goreleaser
	@make dep
	gometalinter --install --update

## Install dependencies of the project
dep:
	@dep ensure -v

## Visualizing dependencies status of the project
dep-status:
	@dep status

dep-add:
	@printf "\n"; \
	read -p "Repository: "; \
	if [ ! "$$REPLY" ]; then \
		printf "\n${COLOR_RED}"; \
		echo "Invalid repository."; \
		exit 1; \
	fi; \
	dep ensure -add $$REPLY 

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
	git commit README.md -m "Update README.md with release $$TAG"; \
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
	git push origin master; \
	echo "Release - OK"


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