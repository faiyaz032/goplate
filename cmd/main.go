package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/faiyaz032/goplate/internal/config"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, cfg); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
