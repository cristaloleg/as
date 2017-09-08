.PHONY: all install build test lint
PKG = $(shell go list ./... | grep -v /vendor/)

all: install build test

install:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/cover

build:
	go build ${PKG}

test:
	go test -v ${PKG}

lint:
	golint ${PKG}
	go vet ${PKG}