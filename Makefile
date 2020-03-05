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
	go get ./...

.PHONY: vet
vet: ## Vet the code
	go vet -v ./...

.PHONY: lint
lint: ## Lint the code
	golint -set_exit_status $(shell go list ./...)

.PHONY: build
build: ## Build the application
	go build .

.PHONY: test
test: ## Run the unit tests
	go test ./...

.PHONY: all ## Run everything
all: install vet lint build test
