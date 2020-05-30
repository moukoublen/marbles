default: build

GO111MODULE?=on
export GO111MODULE

EXEC?=marbles
MAINCMD?=./cmd/marbles

env:
		go env

test:
		go test -race -timeout 60s $$(go list ./...)

test-integration:
		go test -race -v -tags=integration -timeout 60s $$(go list ./...)

test-coverage:
		go test -race -timeout 60s -coverprofile cover.out -covermode atomic $$(go list ./...)
		go tool cover -func cover.out
		rm cover.out

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
