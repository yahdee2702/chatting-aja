package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yahdee2702/chatting-aja/internal/config"
	"github.com/yahdee2702/chatting-aja/internal/logging"
	"github.com/yahdee2702/chatting-aja/internal/server"
)

func main() {
	config := config.Load()

	logger := logging.New(config.App.Env)
	slog.SetDefault(logger)

	server := server.New(config, logger)

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
