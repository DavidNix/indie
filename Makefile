default: help

.PHONY: help
help: ## Print this help message
	@echo "Available make commands:"; grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Setup your local dev environment. Run this once after cloning the repo.
	@# golangci-lint does not recommend using `go get` to install
	@brew install golangci-lint
	brew upgrade golangci-lint
	@mkdir -p .git/hooks
	@cp script/pre-push .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@go get -tool -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: run
run: ## Run the app
	@GOEXPERIMENT=loopvar go run -mod=readonly -race . --addr localhost:3000

.PHONY: vet
vet: ## Run vet and linters
	golangci-lint run

.PHONY: test
test: ## Run unit tests
	@go test -mod=readonly -race -cover -timeout=60s ./...

.PHONY: gen
gen: ## Generate code
	go generate ./...
	go tool goimports -w .
