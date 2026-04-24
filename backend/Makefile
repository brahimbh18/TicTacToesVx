.PHONY: run test build fmt

run:
go run ./cmd/api

test:
go test ./...

build:
go build ./...

fmt:
gofmt -w $(shell find . -name '*.go')
