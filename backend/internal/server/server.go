package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yahdee2702/chatting-aja/internal/config"
)

type Server struct {
	cfg        *config.AppConfig
	logger     *slog.Logger
	httpServer *http.Server
}

func New(config *config.Config, logger *slog.Logger) *Server {
	router := chi.NewRouter()

	return &Server{
		cfg:    &config.App,
		logger: logger,
		httpServer: &http.Server{
			Addr:    ":" + config.App.Port,
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s.logger.Info("starting server", "port", s.cfg.Port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("server shutting down")
	return s.httpServer.Shutdown(ctx)
}
