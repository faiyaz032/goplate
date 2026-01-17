package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/faiyaz032/goplate/internal/auth"
	"github.com/faiyaz032/goplate/internal/config"
	"github.com/faiyaz032/goplate/internal/infrastructure/db/postgres"
	"github.com/faiyaz032/goplate/internal/infrastructure/logger"
	"github.com/faiyaz032/goplate/internal/repository"
	authhandler "github.com/faiyaz032/goplate/internal/rest/handler/auth"
	userhandler "github.com/faiyaz032/goplate/internal/rest/handler/user"
	"github.com/faiyaz032/goplate/internal/rest/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/faiyaz032/goplate/internal/user"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func run(ctx context.Context, cfg *config.Config) error {
	// Logger Setup
	log, err := logger.NewLogger(cfg.AppEnv)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	// infra Setup
	conn, err := postgres.NewConnection(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close(ctx)
	queries := postgres.NewQueries(conn)
	log.Debug("database connected successfully")

	//pkg
	v := validator.New()

	// dependency injection
	userRepo := repository.NewUserRepository(queries, log)
	userSvc := user.NewService(userRepo, log)
	userHndlr := userhandler.NewHandler(userSvc, log)

	authSvc := auth.NewService(userSvc, log)
	authHndlr := authhandler.NewHandler(v, authSvc, log)

	// routing
	r := chi.NewRouter()
	r.Use(middleware.CORS())
	r.Use(middleware.RequestLogger(log))
	r.Use(chiMiddleware.Recoverer)

	userhandler.RegisterRoutes(r, userHndlr)
	authhandler.RegisterRoutes(r, authHndlr)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// server Configuration
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// start server
	go func() {
		log.Info("starting server", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("listen error", zap.Error(err))
		}
	}()

	<-ctx.Done()
	//shutdown
	log.Info("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown failed", zap.Error(err))
		return err
	}

	return nil
}
