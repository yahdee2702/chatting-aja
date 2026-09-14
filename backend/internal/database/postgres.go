package database

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/yahdee2702/chatting-aja/internal/config"
)

func NewPostgres(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := url.URL{
		Scheme: cfg.Scheme,
		User:   url.UserPassword(cfg.User, cfg.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Path:   cfg.Name,
	}

	query := dsn.Query()
	query.Set("sslmode", cfg.SSLMode)

	dsn.RawQuery = query.Encode()
	db, err := sqlx.Open(cfg.Driver, dsn.String())

	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}
