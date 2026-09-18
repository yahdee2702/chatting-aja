package websocket

import (
	"log/slog"
	"net/http"

	"github.com/coder/websocket"
	"github.com/yahdee2702/chatting-aja/internal/chat"
)

type Handler struct {
	logger      *slog.Logger
	chatService *chat.Service
}

func NewHandler(logger *slog.Logger, chatService *chat.Service) *Handler {
	return &Handler{
		logger:      logger,
		chatService: chatService,
	}
}

func (h *Handler) ConnectChat(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		h.logger.Error("failed to accept websocket connection", "error", err)
	}

	defer conn.CloseNow()

	h.logger.Info("websocket connected")
}
