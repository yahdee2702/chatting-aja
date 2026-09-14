package auth

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, authHandler *Handler) {
	r.Post("/login", authHandler.Login)
}
