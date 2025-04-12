default: help

.PHONY: help
help: ## Print this help message
	@echo "Available make commands:"; grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Setup your local dev environment. Run this once after cloning the repo.
	@# golangci-lint does not recommend using `go get` to install
	brew install golangci-lint
	brew upgrade golangci-lint
	@mkdir -p .git/hooks
	cp script/pre-push .git/hooks/pre-push
	chmod +x .git/hooks/pre-push
	go get -tool -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	npm install -D tailwindcss@latest tailwindcss@latest

.PHONY: check-clean
check-clean: ## Check if git state is clean
	@git diff-index --quiet HEAD -- || { echo "Error: Git is dirty."; exit 1; }

FEAT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
.PHONY: pr
pr: check-clean ## Mimic a local PR from a branch into upstream
	@if [ "$(FEAT_BRANCH)" = "main" ]; then \
		echo "Error: You are on the main branch."; \
		exit 1; \
	fi
	@git checkout main
	@git merge --squash $(FEAT_BRANCH)

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

.PHONY: css
css:
	npx @tailwindcss/cli -i ./internal/server/css/app.css -o ./internal/server/static/styles.css --watch

.PHONY: watch
watch: ## Watch and reload code changes
	go tool templ generate --watch --proxy="http://localhost:3000" --cmd='go run -mod=readonly -race ./cmd/indie/... server'

.PHONY: overmind
overmind: ## Run Procfile runner
	 go tool overmind start -f Procfile.dev

GO_PKG := github.com/DavidNix/indie
VERSION := $(shell git describe --tags --always --dirty)

.PHONY: build
build: ## Build local app
	@echo "Building $(VERSION)"
	@mkdir -p release
	go build -tags prod \
    -ldflags "-extldflags=-static -s -X $(GO_PKG)/internal/version.V=$(VERSION)" \
    -o ./release ./cmd/...

.PHONY: build-linux-amd64
build-linux-amd64: ## Cross compile app for linux amd64 using zig
	@echo "Building $(VERSION)"
	@mkdir -p release-linux
	CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    CC="zig cc -target x86_64-linux-musl" \
    go build -tags sqlite_omit_load_extension,prod \
    -ldflags "-extldflags=-static -s -X $(GO_PKG)/internal/version.V=$(VERSION)" \
    -o release-linux ./cmd/...

# DATABASE MIGRATIONS

MIGRATIONS_PATH := internal/database/migrations

.PHONY: db-migrate-new
db-migrate-new: ## Create a new application database migration file. E.g. make db-migrate-new NAME=create_domains_table
	go tool migrate create -ext sql -dir $(MIGRATIONS_PATH) $(NAME)

DB_MIGRATE := go tool migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)"

.PHONY: db-migrate-up
db-migrate-up: ## Migrate the application database up
	@$(DB_MIGRATE) up

.PHONY: db-migrate-force
db-migrate-force: ## Force the application database up to a specific version. E.g. make db-migrate-force VERSION=1
	@$(DB_MIGRATE) force $(VERSION)

.PHONY: db-migrate-down
db-migrate-down: ## Migrate the application database down
	@echo "This is irreversible. Are you sure? (yes/no)"
	@read confirm && if [ $$confirm == "yes" ]; then $(DB_MIGRATE) down 1; else echo "Cancelled" && exit 0; fi