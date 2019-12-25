default: build

GO111MODULE?=on
export GO111MODULE

EXEC?=marbles
MAINCMD?=./cmd/marbles

env:
		go env

test:
		go test $$(go list ./...)

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
		go build -mod=vendor -a -o $(EXEC) $(MAINCMD)

run: vendor
		go run -mod=vendor $(MAINCMD)

clean:
		rm -f $(EXEC)

install:
		cp $(EXEC) /usr/local/bin

uninstall:
		rm -f /usr/local/bin/$(EXEC)

.PHONY: env test fmt lint vet mod vendor build run clean install uninstall
