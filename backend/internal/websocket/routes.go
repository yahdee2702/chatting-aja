package websocket

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router, websocketHandler *Handler) {
	r.Get("/chat", websocketHandler.ConnectChat)
}
