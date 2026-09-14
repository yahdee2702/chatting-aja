package auth

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/yahdee2702/chatting-aja/internal/httpx"
)

type Handler struct {
	service  *Service
	logger   *slog.Logger
	validate *validator.Validate
}

func NewHandler(
	service *Service,
	logger *slog.Logger,
) *Handler {
	return &Handler{
		service:  service,
		logger:   logger,
		validate: validator.New(),
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {
			httpx.ValidationError(w, validationErrors)
			return
		}

		httpx.Error(w, http.StatusBadRequest, "Invalid request")
		return
	}
}
