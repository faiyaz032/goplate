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
   ```bash
   go run ./cmd
   ```

## Project Structure

Refer to the [Architecture Documentation](docs/ARCHITECTURE.md) for details on the project structure and design patterns.

## Makefile Commands

- `make migrate-up`: Apply migrations.
- `make migrate-down`: Rollback last migration.
- `make migration`: Create a new migration file.

## License

MIT
