PLAYGROUND_NAME := playground
PKGS := $(shell go list ./pkg/... ./genproto/...)

all: build

.PHONY: generate playground

generate:
	go generate $(PKGS)
playground:
	go run ./cmd/$(PLAYGROUND_NAME)/.