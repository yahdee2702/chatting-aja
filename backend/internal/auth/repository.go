package auth

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const userColumns = `
    id,
    name,
    email,
    password_hash,
    status,
    email_verified_at,
    created_at,
    updated_at,
    deleted_at
`

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.QueryRowxContext(
		ctx,
		`SELECT `+userColumns+`
		FROM users 
		WHERE email = $1`,
		email,
	).StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FindUserById(ctx context.Context, id string) (*User, error) {
	var user User
	err := r.db.QueryRowxContext(
		ctx,
		`SELECT `+userColumns+`
		FROM users 
		WHERE id = $1`,
		id,
	).StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(ctx context.Context, name string, email string, passwordHash string) (string, error) {
	var id string

	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	if err := r.db.GetContext(ctx, &id, query, name, email, passwordHash); err != nil {
		return "", err
	}

	return id, nil
}
