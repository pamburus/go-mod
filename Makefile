.PHONY: all
all: ci

.PHONY: ci
ci: lint test

.PHONY: lint
lint:
	@go list -m -f '{{.Dir}}' | xargs -I{} golangci-lint run {}/...

.PHONY: test
test:
	@go list -m -f '{{.Dir}}' | xargs -I{} go test {}/...
