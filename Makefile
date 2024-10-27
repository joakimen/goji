BIN = goji
BIN_DIR = bin

all: format lint test build 

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: test
test:
	go test -v ./...

.PHONY: install
install:
	go install

.PHONY: build
build:
	go build -o $(BIN_DIR)/$(BIN) .

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: run
run:
	go run main.go auth --help
