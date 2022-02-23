include common/make/common.mk

SERVICE=sample-app # Ensure the service name matches the name of the repo

## Top level
build: go-build ## Run all builds for this repo
clean: go-clean proto-clean ## Clean up anything extra folders
lint: go-lint proto-lint docker-lint ## Run all available linters
test: go-test ## Run all available unit tests
test-coverage: go-test-cov-html ## Open browser with test coverage tool
test-integration: go-test-integration ## Run all available integration tests
install-tools: go-install-tools alderson-install-tools ## Install all required tools
docker-build: docker-build-app docker-build-integration-tests ## Build all docker images

BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

