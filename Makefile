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

.PHONY: watch
watch: ## Watch and reload code changes
	go tool templ generate --watch --proxy="http://localhost:3000" --cmd='go run -mod=readonly -race .'

.PHONY: overmind
overmind: ## Run Procfile runner
	 go tool overmind start -f Procfile.dev

GO_PKG := github.com/DavidNix/indie
VERSION := $(shell git describe --tags --always --dirty)"

.PHONY: build
build: ## Build local app
	@mkdir -p release
	go build -tags prod \
    -ldflags "-extldflags=-static -s -X $(GO_PKG)/internal/version.Version=$(VERSION) \
    -o ./release ./cmd/...

.PHONY: build-linux-amd64
build-linux-amd64: ## Cross compile app for linux amd64 using zig
	@mkdir -p release-linux
	CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    CC="zig cc -target x86_64-linux-musl" \
    go build -tags sqlite_omit_load_extension,prod \
    -ldflags "-extldflags=-static -s -X $(GO_PKG)/internal/version.Version=$(VERSION)" \
    -o release-linux ./cmd/...
