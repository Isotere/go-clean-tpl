SHELL   := /bin/bash -euo pipefail
TIMEOUT := 10s
GOFLAGS := -mod=vendor
BINPATH = $(PWD)/bin

.PHONY: deps
deps: deps-main deps-tools

.PHONY: deps-main
deps-main:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: deps-tools
deps-tools: