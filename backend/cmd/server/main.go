package server

import (
	"log/slog"

	"github.com/yahdee2702/chatting-aja/internal/config"
	"github.com/yahdee2702/chatting-aja/internal/logging"
	"github.com/yahdee2702/chatting-aja/internal/server"
)

func main() {
	config := config.Load()

	logger := logging.New(config.App.Env)
	slog.SetDefault(logger)

	server := server.New(config, logger)

	if err := server.Run(); err != nil {
		logger.Error("failed to start server", "error", err)
	}
}
