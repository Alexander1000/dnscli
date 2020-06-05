PROJECTNAME := "dnscli"
VERSION := $(shell cat VERSION)
BUILD := $(shell git rev-parse --short HEAD)
DESCRIPTION := "CLI utility for manage DNSaaS"
MAINTAINER := "Mikhail Bruskov <mvbruskov@avito.ru>"
SHELL := /bin/bash

# Go source files, ignore vendor directory
GOFILES = $(shell find . -type f -name '*.go'\
 -not -path "./vendor/*"\
 -not -path "./cmd/*"\
 -not -path "./app/*"\
 -not -path "./models/*"\
 -not -path "./pdnshttp/*")

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X=github.com/mixanemca/dnscli/cmd.version=$(VERSION) -X=github.com/mixanemca/dnscli/cmd.build=$(BUILD)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.SHELLFLAGS = -ec

.PHONY: build help

all: build

## build: Compile the binary.
#build: vet lint test clean
build: clean
	@printf "compile the binary"
	@go build $(LDFLAGS) -o $(PROJECTNAME) $(GOFILES)
	@echo "  ok"

## linux: Compile the binary for GNU/Linux amd64.
linux: vet lint test clean
	@printf "compile the binary for Linux"
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(PROJECTNAME) $(GOFILES)
	@echo "  ok"

## vet: Runs the go vet command on the packages named by the import paths.
vet:
	@printf "runs the go vet command"
	@go vet $(GOFILES)
	@echo "  ok"

## lint: Runs the golint command.
lint:
	@printf "runs the golint command"
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status $(GOFILES)
	@echo "  ok"

## test: Runs the go test command.
test:
	@echo "runs the go test command"
	@go test -v $(GOFILES)

## fpm-deb: Build Debian package.
fpm-deb: linux
	fpm -s dir -t deb -n $(PROJECTNAME) -v $(VERSION) \
		--deb-priority optional --category admin \
		--force \
		--url https://github.com/nemca/$(PROJECTNAME) \
		--description $(DESCRIPTION) \
		-m $(MAINTAINER) \
		--license "Apache 2.0" \
		-a amd64  \
		dnscli.yaml.example=/usr/share/doc/$(PROJECTNAME)/dnscli.yaml.example \
		$(PROJECTNAME)=/usr/bin/$(PROJECTNAME)

## install: Install binary to your system
install: build
	install -d /usr/local/bin
	install -m 0755 $(PROJECTNAME) /usr/local/bin/

## uninstall: Uninstall binary from your system
uninstall:
	@-rm -f /usr/local/bin/$(PROJECTNAME)

## clean: Cleanup.
clean:
	@-rm -f $(PROJECTNAME)
	@-rm -f *.deb

## help: Show this message.
help: Makefile
	@echo "Available targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
