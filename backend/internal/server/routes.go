package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yahdee2702/chatting-aja/internal/auth"
	"github.com/yahdee2702/chatting-aja/internal/chat"
	"github.com/yahdee2702/chatting-aja/internal/websocket"
)

func NewRouter(
	authHandler *auth.Handler,
	chatHandler *chat.Handler,
	websocketHandler *websocket.Handler,
	authMiddleware *auth.AuthMiddleware,
) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			auth.Routes(r, authHandler, authMiddleware)
		})

		r.Route("/chat", func(r chi.Router) {
			r.Use(authMiddleware.Handle)
			chat.Routes(r, chatHandler)
		})
	})

	r.Route("/ws", func(r chi.Router) {
		r.Use(authMiddleware.Handle)
		websocket.Routes(r, websocketHandler)
	})

	return r
}
