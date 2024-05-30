.DEFAULT_GOAL = build

.PHONY: help
help: ## Display help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run linter
	golangci-lint run -v ./...

.PHONY: migrate-up
migrate-up:
	goose -dir ./migrations up

.PHONY: migrate-down
migrate-down:
	goose -dir ./migrations down

.PHONY: migrate-status
migrate-status:
	goose -dir ./migrations status

.PHONY:build
build: ## Prepare binaries
	go build -C ./cmd/client/ -o client
	go build -C ./cmd/server/ -o server

.PHONY: run-client
run-client: ## Run client
	go run ./cmd/client/main.go

.PHONY: run-server
run-server: ## Run server
	go run ./cmd/server/main.go

.PHONY: test
test: ## Run all tests
	go test -v -race ./...

.PHONY: clean
clean: ## Delete binaries
	-rm -f ./cmd/client/client
	-rm -f ./cmd/server/server
