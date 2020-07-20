DEV_NETWORK = megaphone
POSTGRES_VERSION = latest
DOCKER_COMPOSE_ARGS = --file ./deployments/docker-compose.yml \
					  --project-name sales_admin \
					  --project-directory .


help: ## print help message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%s\033[0m| %s\n", $$1, $$2}' | column -s '|' -t

postgres-up: ## Start postgres container to test against during development.
	@POSTGRES_VERSION=latest \
	POSTGRES_PASSWORD=test \
	docker-compose $(DOCKER_COMPOSE_ARGS) up -d postgres migrate


postgres-down: ## Stop the development postgres container.
	@POSTGRES_VERSION=latest \
	POSTGRES_PASSWORD=test \
	docker-compose $(DOCKER_COMPOSE_ARGS) down postgres migrate

server: ## Build the sale_admin server.
	@GO111MODULES=on go build ./cmd/server

.PHONY: test
test: postgres-up ## Run the tests for the application.
	@docker run \
		--rm \
		-ti \
		--network salesadmin_megaphone \
		-v $$(pwd):/app \
		-w /app \
		golang:latest \
		go test ./...

up: ## Run the sales_admin server and all it's dependencies.
	@POSTGRES_VERSION=$(POSTGRES_VERSION) \
	POSTGRES_PASSWORD=test \
	docker-compose $(DOCKER_COMPOSE_ARGS) up --force-recreate --build -d



tail-logs: ## Tail the logs of the running docker-compose stack.
	@POSTGRES_VERSION=$(POSTGRES_VERSION) \
	POSTGRES_PASSWORD=test \
	docker-compose $(DOCKER_COMPOSE_ARGS) logs -f

down: ## stop the running sales_admin server and its dependencies.
	@POSTGRES_VERSION=$(POSTGRES_VERSION) \
	POSTGRES_PASSWORD=test \
	docker-compose $(DOCKER_COMPOSE_ARGS) down


