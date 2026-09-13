package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yahdee2702/chatting-aja/internal/config"
)

type Server struct {
	cfg    *config.AppConfig
	logger *slog.Logger
}

func New(config *config.Config, logger *slog.Logger) *Server {
	return &Server{
		cfg:    &config.App,
		logger: logger,
	}
}

func (srv *Server) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return http.ListenAndServe(fmt.Sprintf(":%s", srv.cfg.Port), r)
}
