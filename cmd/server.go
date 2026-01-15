package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/faiyaz032/goplate/internal/auth"
	"github.com/faiyaz032/goplate/internal/config"
	"github.com/faiyaz032/goplate/internal/infrastructure/db/postgres"
	"github.com/faiyaz032/goplate/internal/repository"
	authHandler "github.com/faiyaz032/goplate/internal/rest/auth"
	userHandler "github.com/faiyaz032/goplate/internal/rest/user"
	"github.com/faiyaz032/goplate/internal/user"
	"github.com/go-chi/chi/v5"
)

func run(ctx context.Context, cfg *config.Config) error {
	// infra Setup
	conn, err := postgres.NewConnection(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close(ctx)
	queries := postgres.NewQueries(conn)

	// dependency injection
	userRepo := repository.NewUserRepository(queries)
	userSvc := user.NewService(userRepo)
	userHndlr := userHandler.NewHandler(userSvc)

	authSvc := auth.NewService(userSvc)
	authHndlr := authHandler.NewHandler(authSvc)

	// routing
	r := chi.NewRouter()
	userHandler.RegisterRoutes(r, userHndlr)
	authHandler.RegisterRoutes(r, authHndlr)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// server Configuration
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// start server
	go func() {
		fmt.Printf("Server running on http://localhost:8080\n")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Listen error: %s\n", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("\nShutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(shutdownCtx)
}
