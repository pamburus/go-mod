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
all-modules := $(subst \,/,$(all-modules))
all-modules := $(all-modules:$(abspath .)/%=%)
endif

# Auxiliary modules, not to be tested
aux-modules += build/tools

# Populate module list to test
ifndef modules
modules := $(filter-out $(aux-modules),$(all-modules))
endif

# Tools
go-test := go test
go-tool-cover := go tool cover
coverage-filter := go run ./build/tools/cmd/coverage-filter
test-filter := go run ./build/tools/cmd/test-filter
ifeq ($(verbose),yes)
	coverage-filter += -v
	go-test += -v
	export SQLTEST_LOG_LEVEL=debug
endif
ifeq ($(debug),yes)
	export SQLTEST_DEBUG=yes
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
ifeq ($(debug),yes)
	go list -f '{{.Dir}}' ./$* | xargs -o -S 4096 -I {} ${SHELL} -c 'go test -c -o {}/.test -coverprofile={}/.cover.out {} && {}/.test -test.v -test.coverprofile={}/.cover.out'
else
	$(go-test) -fullpath -coverprofile=$*/.cover.out ./$*/... | $(test-filter)
endif

# ---

## Show coverage
.PHONY: coverage
coverage: $(modules:%=coverage/%)

## Show coverage for a module
.PHONY: coverage/%
coverage/%: test/%
	$(go-tool-cover) -func=$*/.cover.out | $(coverage-filter) $(import-path)

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
	find . -type f -name .cover.out -delete
	find . -type f -name .test -delete
	find . -type f -name go.work.sum -delete
