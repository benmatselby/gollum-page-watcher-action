.PHONY: explain
explain:
	### Welcome
	#
	# $$ make all
	#
	### Targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## Install the local dependencies
	go install github.com/golang/mock/mockgen@master
	go get ./...

.PHONY: lint
lint: ## Lint the code
	golangci-lint run

.PHONY: build
build: ## Build the application
	go build .

.PHONY: test
test: ## Run the unit tests
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: test-cov
test-cov: test ## Run the unit tests with coverage
	go tool cover -html=coverage.out

.PHONY: all ## Run everything
all: install lint build test
