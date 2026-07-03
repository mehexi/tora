APP     := tora
CMD     := ./cmd/tora
BIN     := ./bin/$(APP)
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

.PHONY: all build run clean test lint fmt vet install

all: build

build:
	@mkdir -p bin
	go build $(LDFLAGS) -o $(BIN) $(CMD)

run:
	go run $(CMD)

install:
	go install $(LDFLAGS) $(CMD)

test:
	go test ./...

test-v:
	go test -v ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

clean:
	rm -rf bin/

.DEFAULT_GOAL := build
