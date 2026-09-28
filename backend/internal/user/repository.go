package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUsersWithLimit(ctx context.Context, limit int) ([]User, error) {
	var users []User
	if err := r.db.SelectContext(
		ctx,
		&users,
		`SELECT id, name, email, status, created_at
		FROM users
		WHERE deleted_at IS NULL
		LIMIT $1`,
		limit,
	); err != nil {
		return nil, err
	}

	return users, nil
}
