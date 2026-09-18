package chat

import (
	"log/slog"
)

type Handler struct {
	logger  *slog.Logger
	service *Service
}

func NewHandler(logger *slog.Logger, service *Service) *Handler {
	return &Handler{
		logger: logger,
	}
}
