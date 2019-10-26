.PHONY: clean mod vendor build install uninstall

default: build

mod:
		go mod tidy
		go mod verify

vendor: mod
		go mod vendor

build: vendor
		GO111MODULE=on go build -mod=vendor -a -o marbles ./cmd/marbles

clean:
		rm -f marbles

install:
		cp marbles /usr/local/bin

uninstall:
		rm -f /usr/local/bin
