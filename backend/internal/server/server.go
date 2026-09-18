package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/yahdee2702/chatting-aja/internal/auth"
	"github.com/yahdee2702/chatting-aja/internal/chat"
	"github.com/yahdee2702/chatting-aja/internal/config"
	"github.com/yahdee2702/chatting-aja/internal/websocket"
)

type Server struct {
	cfg        *config.AppConfig
	logger     *slog.Logger
	db         *sqlx.DB
	httpServer *http.Server
}

func New(config *config.Config, logger *slog.Logger, db *sqlx.DB) *Server {
	jwtHandler := auth.NewJwtHandler([]byte(config.App.Secret))

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository, jwtHandler)
	authHandler := auth.NewHandler(authService, logger)

	chatRepository := chat.NewRepository(db)
	chatService := chat.NewService(chatRepository)
	chatHandler := chat.NewHandler(logger, chatService)

	websocketHandler := websocket.NewHandler(logger, chatService)

	authMiddleware := auth.NewAuthMiddleware(jwtHandler)

	router := NewRouter(authHandler, chatHandler, websocketHandler, authMiddleware)

	return &Server{
		cfg:    &config.App,
		logger: logger,
		httpServer: &http.Server{
			Addr:    ":" + config.App.Port,
			Handler: router,
		},
		db: db,
	}
}

func (s *Server) Run() error {
	s.logger.Info("starting server", "port", s.cfg.Port)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("server shutting down")
	return s.httpServer.Shutdown(ctx)
}
