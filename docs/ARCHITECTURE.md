# Architecture Documentation

This document describes the architectural patterns and folder structure used in the `goplate` boilerplate.

## Architectural Patterns

The project follows a **Layered Architecture** with hints of **Clean Architecture** and **Hexagonal Architecture** principles. This ensures a separation of concerns, making the codebase maintainable and testable.

### Core Principles

- **Dependency Inversion**: High-level modules do not depend on low-level modules. Both depend on abstractions.
- **Dependency Injection**: Dependencies are injected into components (manual DI is performed in `cmd/server.go`).
- **Separation of Layers**: Each layer has a specific responsibility.

## Folder Structure

```text
.
├── cmd/                   # Main applications (entry points)
│   ├── main.go            # Application bootstrap
│   └── server.go          # Server setup and Dependency Injection
├── internal/              # Private application code
│   ├── auth/              # Authentication service layer
│   ├── config/            # Configuration management
│   ├── domain/            # Domain entities and core logic (lowest layer)
│   ├── infrastructure/    # External tools (DB connections, etc.)
│   ├── repository/        # Data access layer (SQLC implementations)
│   ├── rest/              # HTTP delivery layer
│   │   ├── handler/       # HTTP request handlers
│   │   └── middleware/    # Custom HTTP middlewares
│   └── user/              # User service layer
├── migrations/            # Database migrations (Goose)
├── queries/               # SQL queries for SQLC
└── sqlc.yml               # SQLC configuration
```

## Layers

### 1. Domain (`internal/domain`)
The innermost layer. Contains business entities (structs) and domain-specific errors. It has no dependencies on other internal layers.

### 2. Service Layer (`internal/user`, `internal/auth`)
Contains business logic and use cases. It interacts with the repository layer through abstractions.

### 3. Repository Layer (`internal/repository`)
Responsible for data persistence. It uses the generated code from `sqlc` to interact with PostgreSQL via `pgx`.

### 4. REST Layer (`internal/rest`)
The delivery layer.
- **Handlers**: Orchestrate request parsing, service calls, and response formatting.
- **Middleware**: Cross-cutting concerns like CORS and logging.

### 5. Infrastructure Layer (`internal/infrastructure`)
Contains implementations for external resources like database connections.

## Data Flow

1. **Request**: A client sends an HTTP request.
2. **REST Layer**: The router (`chi`) directs the request to a **Handler**.
3. **Validation**: The handler validates the input using `validator/v10`.
4. **Service Call**: The handler calls a method on the **Service**.
5. **Business Logic**: The service performs logic and calls the **Repository** if needed.
6. **Persistence**: The repository executes SQL queries using generated `sqlc` code.
7. **Response**: The results bubble back up and are returned as JSON.

## Dependency Injection

Dependency Injection is handled manually in `cmd/server.go`. This provides a clear overview of how the application is wired together without the magic of reflection-based DI containers.
