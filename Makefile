GO				:= $(shell which go)
RM				:= rm -f
SHELL			:= /bin/bash

CGO_ENABLED	:= 0
GOFILES		= $(shell go list ./... | grep -v /vendor/)

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) build cmd/ternakgopher/ternakgopher.go

.PHONY: all
all: install lint test build

.PHONY: clean
clean:
	$(RM) ternakgopher

.PHONY: install
install:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint
	$(GOPATH)/bin/dep ensure -v

.PHONY: lint
lint:
	@$(GOPATH)/bin/golint $(GOFILES)

.PHONY: test
test:
	$(GO) test ./...