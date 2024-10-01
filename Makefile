# Common makefile helpers
include build/make/common.mk

# Set default goal
.DEFAULT_GOAL := all

# Populate module list
ifndef modules
modules := $(shell go list -m -f '{{.Dir}}')
modules := $(modules:$(PWD)/%=%)
endif

## Run all tests
all: ci
.PHONY: all

# ---

## Run continuous integration tests
ci: lint test
.PHONY: ci

## Run continuous integration tests for a module
ci/%: lint/% test/%
	@true
.PHONY: ci/%

# ---

## Run linters
lint: $(modules:%=lint/%)
.PHONY: lint

## Run linters for a module
lint/%:
	golangci-lint run ./$*/...
.PHONY: lint/%

# ---

## Run tests
test: $(modules:%=test/%)
.PHONY: test

## Run tests for a module
test/%:
	go test -coverprofile=./$*/.cover.out ./$*/...
.PHONY: test/%

# ---

## Show coverage
coverage: $(modules:%=coverage/%)
.PHONY: coverage

## Show coverage for a module
coverage/%: test/%
	go tool cover -func=./$*/.cover.out
.PHONY: coverage/%

# ---

## Clean up
clean:
	rm -f $(modules:%=%/.cover.out)
.PHONY: clean
