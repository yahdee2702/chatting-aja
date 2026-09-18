package auth

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router, authHandler *Handler, authMiddleware *AuthMiddleware) {
	r.Post("/login", authHandler.Login)
	r.Post("/register", authHandler.Register)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.Handle)

		r.Get("/me", authHandler.Me)
	})
}
