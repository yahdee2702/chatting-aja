package chat

import (
	"log/slog"
	"net/http"

	"github.com/coder/websocket"
)

type Handler struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		h.logger.Error("failed to accept websocket connection", "error", err)
	}

	defer conn.CloseNow()

	h.logger.Info("websocket connected")
}
