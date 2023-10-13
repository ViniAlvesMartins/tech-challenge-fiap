GOPATH ?= $(shell go env GOPATH)
GO_RELOAD = $(GOPATH)/bin/reflex -s -r '\.go$$' --
GO_TAGS := musl

ifeq ($(shell uname && uname -p), Darwin arm)
	GO_TAGS := dynamic
endif

install:
	@go install github.com/cespare/reflex@latest

run-api:
	@$(GO_RELOAD) go run -tags $(GO_TAGS) -race main.go
