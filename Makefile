SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: lint
lint: ## Run linters.
	@golangci-lint version
	golangci-lint run ./...

.PHONY: test
test: ## Run tests.
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: test-coverage
test-coverage: test ## Run tests and open coverage report in browser.
	go tool cover -func coverage.txt | grep total
	go tool cover -html=coverage.txt
