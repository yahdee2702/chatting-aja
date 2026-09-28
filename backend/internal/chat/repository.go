package chat

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

func (r *Repository) FindChatByIds(ctx context.Context, userId string, targetUserId string) (*Chat, error) {
	var chat Chat
	row := r.db.QueryRowxContext(ctx, `
	SELECT
		c.id,
		c.type,
		c.created_at
	FROM chats c
	JOIN chats_users cu ON cu.chat_id = c.id
	WHERE c.type = 'individual'
		AND cu.user_id IN ($1, $2)
	GROUP BY c.id
	HAVING COUNT(DISTINCT cu.user_id) = 2;
	`)

	if err := row.StructScan(&chat); err != nil {
		return nil, err
	}

	return &chat, nil
}
