GOPATH := $(PWD)
PATH := "$(PATH):$(GOPATH)/bin:/opt/golang/bin"

.PHONY: all test

all: build test lint

test: ## Run unit tests
	go test ./...

