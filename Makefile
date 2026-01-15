# Load environment variables from .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Configuration
MIGRATIONS_DIR=migrations
DOCKER_CONFIG=docker-compose.local.yml

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# --- Docker Commands ---

.PHONY: up
up: ## Start the database using the local compose file
	docker compose -f $(DOCKER_CONFIG) up -d --remove-orphans

# --- Database Migrations (Goose) ---

.PHONY: migrate-status
migrate-status: ## Show the status of all migrations
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status

.PHONY: migrate-up
migrate-up: ## Run all pending migrations
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down: ## Roll back the last migration
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

.PHONY: migration
migration: ## Create a new migration file (Usage: make migration)
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql
