#!/usr/bin/make

.DEFAULT_GOAL := build

GOOS := linux
GOARCH := amd64
LD_FLAGS := -ldflags "-X main.Version=`git describe --tags` -X main.BuildDate=`date -u +%Y-%m-%d_%H:%M:%S` -X main.GitCommit=`git rev-parse HEAD`"

.PHONY: build
build:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LD_FLAGS) -o "dist/secrets-init_$(GOOS)-$(GOARCH)"
	@chmod +x dist/secrets-init_linux-amd64

.PHONY: bin
bin:
	@go build $(LD_FLAGS)

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test ./... -v -race

.PHONY: all
all:
	@make -s build bin test lint

