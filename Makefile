.PHONY: all
all: test lint build

.PHONY: build
build:
	go build

.PHONY: fmt
fmt:
	gofumpt -l -w .

.PHONY: lint
lint: build
	golangci-lint run --enable revive --enable gofumpt --enable godot --enable errname

.PHONY: test
test:
	go test ./...
