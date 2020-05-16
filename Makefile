GOPATH := $(PWD)
PATH := "$(PATH):$(GOPATH)/bin:/opt/golang/bin"
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v /pkg/ | grep -v _test.go)
PKG_LIST := $(shell go list merchant/${PKG}/... | grep -v /vendor/)

.PHONY: all prepare test lint build

all: build test lint

prepare: ## Download and install tools
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(GOPATH)/bin v1.19.1
	@go get -v github.com/beego/bee
	@go get -v github.com/golang/dep/cmd/dep
	@go get -v github.com/jstemmer/go-junit-report
	@go get -v github.com/vektra/mockery/.../
	@go get -u golang.org/x/lint/golint

prepare-migrate:
	@go get -v github.com/beego/bee
	@go get -v github.com/golang/dep/cmd/dep

lint: ## Lint the files
	@./scripts/golangci-lint.sh $(SUB_PROJECTS)
	@./scripts/lint.sh $(SUB_PROJECTS)

build:
	@for project in $(SUB_PROJECTS) ; do \
		echo "Building $$project" && cd src/$$project && make build && cd - ; \
	done

test: ## Run unit tests
	@./scripts/unittest.sh $(SUB_PROJECTS)

dep: ## Get the dependencies
	@cd src/merchant && dep ensure -v

fmt: ## Format all *.go package
	@cd src/ && go fmt $(PKG_LIST) && cd -
