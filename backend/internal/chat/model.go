package chat

import "time"

type Chat struct {
	Id        string    `db:"id"`
	Type      ChatType  `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}
