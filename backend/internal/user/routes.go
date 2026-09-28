package user

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router, userHandler *Handler) {
	r.Get("/", userHandler.GetUsers)
}
