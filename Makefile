all: install build test

install:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/cover

build:
	go build ${PKG}

test:
	go test -v ${PKG}