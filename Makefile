# Makefile for leetcode-go

.DEFAULT_GOAL := all

.PHONY: help
help: ## Display this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: fmt lint tests

.PHONY: fmt
fmt: ## Format code using golangci-lint
	@echo "Formatting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint fmt; \
	else \
		echo "golangci-lint not found. Install it as described in documentation https://golangci-lint.run/docs/"; \
		exit 1; \
	fi

.PHONY: lint
lint: ## Lint code with golangci-lint
	@echo "Running golangci-lint..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install it as described in documentation https://golangci-lint.run/docs/"; \
		exit 1; \
	fi

.PHONY: lint-fix
lint-fix: ## Run golangci-lint with auto-fix
	@echo "Running golangci-lint with auto-fix..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run --fix; \
	else \
		echo "golangci-lint not found. Install it as described in documentation https://golangci-lint.run/docs/"; \
		exit 1; \
	fi

.PHONY: tests
tests: ## Run tests
	@echo "Runnning tests..."
	@go test ./...
