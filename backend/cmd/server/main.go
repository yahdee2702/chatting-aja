package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yahdee2702/chatting-aja/internal/config"
	"github.com/yahdee2702/chatting-aja/internal/database"
	"github.com/yahdee2702/chatting-aja/internal/logging"
	"github.com/yahdee2702/chatting-aja/internal/server"
)

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	logger := logging.New(config.App.Env)
	slog.SetDefault(logger)

	db, err := database.NewPostgres(config.DB)
	if err != nil {
		logger.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	server := server.New(config, logger, db)

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- server.Run()
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case err := <-serverErr:
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	case sig := <-signalChan:
		logger.Info("shutdown signal received", "signal", sig)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server", "error", err)
		os.Exit(1)
	}

	logger.Info("server stopped")
}
