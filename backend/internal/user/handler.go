package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yahdee2702/chatting-aja/internal/httpx"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitStr := query.Get("limit")

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		limit = -1
	}

	users, err := h.service.GetUsers(r.Context(), limit)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, fmt.Sprintf("Cannot fetch users: %s", err.Error()))
		return
	}

	httpx.Success(w, http.StatusOK, "Successfully fetching users", users)
}
