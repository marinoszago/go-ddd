DOCKER_COMPOSE_YAML := deployment/docker-compose.yaml
MAIN_FILE_DIR := cmd/api/main.go
MAIN_IMPORT_FILE_DIR := cmd/import/main.go

lint: ## Perform linting
	golangci-lint run

run: ## Run application
	go run $(MAIN_FILE_DIR)

run_import: ## Run import application
	go run $(MAIN_IMPORT_FILE_DIR)

database: ## Start a mysql container
	docker-compose -f $(DOCKER_COMPOSE_YAML) up -d

gen: ## Generate stuff like mocks
	go generate ./...
	go mod tidy

test: ## Run tests
	go clean -cache
	go test `go list ./... | grep -v 'mock' | grep -v 'cmd' | grep -v 'config'` -race

.PHONY: lint run run_import database gen test