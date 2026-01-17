

############################
# Base (shared)
############################
FROM golang:1.25-alpine AS base

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

############################
# Dev target (Air)
############################
FROM base AS dev

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

############################
# Builder (Prod build)
############################
FROM base AS builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

############################
# Prod runtime
############################
FROM alpine:latest AS prod

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
