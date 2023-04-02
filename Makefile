default: clean run-server tests coverage build k6

BINARY_NAME=server


.PHONY: help
help: ## Print this help with list of available commands/targets and their purpose
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: clean
clean:  ## cleanup test cover output
	@echo "Cleaning up..."
	@docker-compose down
	@rm -f $(BINARY_NAME)

.PHONY: run-server
run-server: ## Run the Go REST API and its dependencies in Docker containers
	@echo "Starting the API server ..."
	@docker-compose up -d --build

.PHONY: tests
tests:  ## run unit tests
	@echo "Running tests..."
	@go test -v -race -count=1 $(shell go list ./... | grep -v pkg)

.PHONY: coverage
coverage:  ## run unit tests
	@echo "Running test coverage ..."
	@go test -coverprofile=coverage.out $(shell go list ./... | grep -v pkg)
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: build
build:  ## build the server
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o $(BINARY_NAME)

.PHONY: k6
k6:  ## Simulate multiple requests in the same time using k6
	k6 run k6.js
