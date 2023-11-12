default: help

.PHONY: help
help: ## Print this help message
	@echo "Available make commands:"; grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Setup your local dev environment. Run this once after cloning the repo.
	@# golangci-lint does not recommend using `go get` to install
	brew install golangci-lint
	brew upgrade golangci-lint

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
	@go generate ./...

AIR = go run -mod=readonly github.com/cosmtrek/air
.PHONY: watch
watch: ## Watch and reload code changes
	@$(AIR)

.PHONY: ent
ent: ## Run ent codegen. E.g. make ent new User
	@go run -mod=readonly entgo.io/ent/cmd/ent $(filter-out $@,$(MAKECMDGOALS))

%: # Catch-all target to allow passing arguments to targets without workarounds like ARGS="1 2 3"
	@:
