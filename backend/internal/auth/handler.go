package auth

import (
	"encoding/json"
	"errors"
	"fmt"
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
		httpx.Error(w, http.StatusBadRequest, "Invalid request body")
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

	token, err := h.service.Login(r.Context(), req)

	if err != nil {
		httpx.Error(w, http.StatusUnauthorized, fmt.Sprintf("Cannot login: %s", err.Error()))
		return
	}

	httpx.Success(w, http.StatusOK, "Succesfully logged in", LoginResponse{
		Token: token,
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid request body")
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

	token, err := h.service.Register(r.Context(), req)

	if err != nil {
		httpx.Error(w, http.StatusUnauthorized, fmt.Sprintf("Cannot register: %s", err.Error()))
		return
	}

	httpx.Success(w, http.StatusCreated, "Succesfully register in", RegisterResponse{
		Token: token,
	})
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)

	user, err := h.service.Me(r.Context(), userId)

	if err != nil {
		httpx.Error(w, http.StatusNotFound, "User not found")
		return
	}

	httpx.Success(w, http.StatusOK, "User found", MeResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}
