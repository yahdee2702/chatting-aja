package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yahdee2702/chatting-aja/internal/auth"
)

func NewRouter(
	authHandler *auth.Handler,
) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			auth.Routes(r, authHandler)
		})
	})

	return r
}
