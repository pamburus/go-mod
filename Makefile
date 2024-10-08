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
	go test -coverprofile=$*/.cover.out ./$*/... | awk -v str="$(PWD)/" '{gsub(str, ""); print}'

# ---

## Show coverage
.PHONY: coverage
coverage: $(modules:%=coverage/%)

## Show coverage for a module
.PHONY: coverage/%
coverage/%: test/%
	go tool cover -func=$*/.cover.out | sed 's|^github\.com/pamburus/go-mod/||' | column -t

# ---

## Tidy up
.PHONY: tidy
tidy: $(modules:%=tidy/%)
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
