SHELL := /bin/bash

NAME := marbles
PKG := github.com/moukoublen/${NAME}
MAINCMD := ./cmd/${NAME}

GO111MODULE := on
export GO111MODULE
CGO_ENABLED := 0

IMAGE := marbles:latest
VERSION?=0.0.1
VER_FLAGS=-X ${PKG}/version=${VERSION}
GO_LDFLAGS=-ldflags "-w -s ${VER_FLAGS}"
GO_LDFLAGS_STATIC=-ldflags "-w -s ${VER_FLAGS} -extldflags -static"

#Commands
GO := go
DOCKER := docker
COMPOSE := docker-compose

.PHONY: default
default: static

.PHONY: env
env:
		$(GO) env

.PHONY: test
test:
		$(GO) test -timeout 60s $(shell go list ./... | grep -v vendor)

.PHONY: test-coverage
test-coverage:
		$(GO) test -timeout 60s -coverprofile cover.out -covermode atomic $(shell go list ./... | grep -v vendor)
		$(GO) tool cover -func cover.out
		rm cover.out

.PHONY: fmt
fmt: ## Verifies all files have been `gofmt`ed.
	@echo "+ $@"
	@if [[ ! -z "$(shell gofmt -s -l . | grep -v vendor | tee /dev/stderr)" ]]; then \
		exit 1; \
	fi

.PHONY: lint
lint:
		$(shell golint ./... | grep -v vendor)

lint-ci:
		$(DOCKER) run --rm -v $(shell pwd):/app:ro -w /app golangci/golangci-lint:v1.27.0 golangci-lint run

.PHONY: vet
vet:
		$(GO) vet $(shell $(GO) list ./...| grep -v vendor)

.PHONY: mod
mod:
		$(GO) mod tidy
		$(GO) mod verify

.PHONY: vendor
vendor:
		$(GO) mod vendor

.PHONY: build
build:
		$(GO) build -a -mod=vendor ${GO_LDFLAGS} -o ${NAME} ${MAINCMD}

.PHONY: static
static:
		CGO_ENABLED=${CGO_ENABLED} $(GO) build -a -mod=vendor ${GO_LDFLAGS_STATIC} -o ${NAME} ${MAINCMD}

.PHONY: clean
clean:
		rm -f ${NAME}

.PHONY: install
install:
		cp ${NAME} /usr/local/bin

.PHONY: uninstall
uninstall:
		rm -f /usr/local/bin/${NAME}
