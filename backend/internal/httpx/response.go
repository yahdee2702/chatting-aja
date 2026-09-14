package httpx

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func JSON(w http.ResponseWriter, status int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(response)
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	JSON(
		w,
		status,
		Response{
			Status:  true,
			Message: message,
			Data:    data,
		},
	)
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(
		w,
		status,
		Response{
			Status:  false,
			Message: message,
		},
	)
}

func ValidationError(w http.ResponseWriter, errors any) {
	JSON(
		w,
		http.StatusBadRequest,
		Response{
			Status:  true,
			Message: "Validation failed",
			Errors:  errors,
		},
	)
}
