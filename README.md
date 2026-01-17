# goplate

A Go boilerplate for building RESTful APIs.

## Tech Stack

| Category | Library |
| :--- | :--- |
| **Router** | [chi/v5](https://github.com/go-chi/chi) |
| **Database** | PostgreSQL |
| **DB Driver** | [pgx/v5](https://github.com/jackc/pgx) |
| **SQL Generation**| [sqlc](https://sqlc.dev/) |
| **Migrations** | [goose](https://github.com/pressly/goose) |
| **Validation** | [validator/v10](https://github.com/go-playground/validator) |
| **Config** | [env/v11](https://github.com/caarlos0/env) |
| **Logging** | [zap](https://github.com/uber-go/zap) |
| **Hot Reload** | [air](https://github.com/air-verse/air) |

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose
- sqlc
- goose

### Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/faiyaz032/goplate.git
   cd goplate
   ```

2. **Setup environment**:
   ```bash
   cp .env.example .env
   ```

3. **Run migrations**:
   ```bash
   make migrate-up
   ```

4. **Run the application**:

   **For Development (Docker with Hot Reload):**
   ```bash
   make dev
   ```

   **For Production (Docker Build):**
   ```bash
   make up
   ```

   **Standalone (Local):**
   ```bash
   go run ./cmd
   ```

## Project Structure

Refer to the [Architecture Documentation](docs/ARCHITECTURE.md) for details on the project structure and design patterns.

## Makefile Commands

- `make dev`: Start development environment with **Docker hot reload** (via Air).
- `make up`: Start production environment.
- `make down`: Stop and remove containers.
- `make migrate-up`: Apply migrations.
- `make migrate-down`: Rollback last migration.
- `make migration`: Create a new migration file.

## License

MIT
