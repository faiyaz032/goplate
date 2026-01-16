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
	authhandler "github.com/faiyaz032/goplate/internal/rest/handler/auth"
	userhandler "github.com/faiyaz032/goplate/internal/rest/handler/user"
	"github.com/go-playground/validator/v10"

	"github.com/faiyaz032/goplate/internal/rest/middleware"
	"github.com/faiyaz032/goplate/internal/user"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func run(ctx context.Context, cfg *config.Config) error {
	// infra Setup
	conn, err := postgres.NewConnection(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close(ctx)
	queries := postgres.NewQueries(conn)

	//pkg
	v := validator.New()

	// dependency injection
	userRepo := repository.NewUserRepository(queries)
	userSvc := user.NewService(userRepo)
	userHndlr := userhandler.NewHandler(userSvc)

	authSvc := auth.NewService(userSvc)
	authHndlr := authhandler.NewHandler(v, authSvc)

	// routing
	r := chi.NewRouter()
	r.Use(middleware.CORS())
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	userhandler.RegisterRoutes(r, userHndlr)
	authhandler.RegisterRoutes(r, authHndlr)

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
