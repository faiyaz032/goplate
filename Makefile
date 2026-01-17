# Load environment variables from .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# --------------------------------------------------
# Configuration
# --------------------------------------------------
MIGRATIONS_DIR   = migrations
COMPOSE_FILE     = docker-compose.yml
COMPOSE_DEV_FILE = docker-compose.dev.yml

.DEFAULT_GOAL := help

# --------------------------------------------------
# Help
# --------------------------------------------------
.PHONY: help
help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# --------------------------------------------------
# Docker
# --------------------------------------------------
.PHONY: dev
dev: ## Start development environment (with hot reload)
	docker compose -f $(COMPOSE_DEV_FILE) up -d --build
	docker compose -f $(COMPOSE_DEV_FILE) logs -f app

.PHONY: up
up: ## Start production environment
	docker compose -f $(COMPOSE_FILE) up -d --build --remove-orphans

.PHONY: down
down: ## Stop and remove all containers
	docker compose -f $(COMPOSE_FILE) down --remove-orphans
	docker compose -f $(COMPOSE_DEV_FILE) down --remove-orphans

.PHONY: restart
restart: ## Restart dev environment
	docker compose -f $(COMPOSE_DEV_FILE) restart

.PHONY: logs
logs: ## Follow dev container logs
	docker compose -f $(COMPOSE_DEV_FILE) logs -f app

.PHONY: ps
ps: ## Show running containers
	docker compose -f $(COMPOSE_DEV_FILE) ps

# --------------------------------------------------
# Database Migrations
# --------------------------------------------------
.PHONY: migrate-status
migrate-status: ## Show migration status
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status

.PHONY: migrate-up
migrate-up: ## Run migrations
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down: ## Roll back last migration
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

.PHONY: migrate-reset
migrate-reset: ## Reset and re-run migrations
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" reset
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

.PHONY: migration
migration: ## Create a new migration file
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql
