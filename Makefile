# Common makefile helpers
include build/make/common.mk

# Common configuration
SHELL := $(SHELL) -o pipefail

# Set default goal
.DEFAULT_GOAL := all

# Some constants
import-path := github.com/pamburus/go-mod

# Populate complete module list, including build tools
ifndef all-modules
all-modules := $(shell go list -m -f '{{.Dir}}')
all-modules := $(all-modules:$(PWD)/%=%)
endif

# Populate module list to test
ifndef modules
modules := $(filter-out build/tools,$(all-modules))
endif

## Run all tests
.PHONY: all
all: ci

# ---

## Run continuous integration tests
.PHONY: ci
ci: lint test

## Run continuous integration tests for a module
.PHONY: ci/%
ci/%: lint/% test/%
	@true

# ---

## Run linters
.PHONY: lint
lint: $(modules:%=lint/%)

## Run linters for a module
.PHONY: lint/%
lint/%:
	golangci-lint run $*/...

# ---

## Run tests
.PHONY: test
test: $(modules:%=test/%)

## Run tests for a module
.PHONY: test/%
test/%:
	go test $(if $(verbose),-v) -coverprofile=$*/.cover.out ./$*/... | go run ./build/tools/test-filter

# ---

## Show coverage
.PHONY: coverage
coverage: $(modules:%=coverage/%)

## Show coverage for a module
.PHONY: coverage/%
coverage/%: test/%
	go tool cover -func=$*/.cover.out | go run ./build/tools/coverage-filter ${import-path} | column -t

# ---

## Tidy up
.PHONY: tidy
tidy: $(all-modules:%=tidy/%)
	go work sync

## Tidy up a module
.PHONY: tidy/%
tidy/%:
	cd $* && go mod tidy

# ---

## Clean up
.PHONY: clean
clean:
	rm -f $(modules:%=%/.cover.out)
