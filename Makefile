.PHONY: build run test lint

build:
	go build -o bin/exchange

run: build
	./bin/exchange

test:
	go test -v ./...

lint:
	go vet .
	golint ./... || true
	staticcheck ./... || true
