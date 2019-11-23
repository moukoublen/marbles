default: build

GO111MODULE?=on
export GO111MODULE

env:
		go env

fmt:
		go fmt $$(go list ./...)

lint:
		golint $$(go list ./...)

vet:
		go vet $$(go list ./...)

mod:
		go mod tidy
		go mod verify

vendor: mod
		go mod vendor

build: vendor
		go build -mod=vendor -a -o marbles ./cmd/marbles

run: vendor
		go run -mod=vendor ./cmd/marbles

clean:
		rm -f marbles

install:
		cp marbles /usr/local/bin

uninstall:
		rm -f /usr/local/bin

.PHONY: env fmt lint vet mod vendor build run clean install uninstall
