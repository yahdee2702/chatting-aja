package auth

import "time"

type User struct {
	Id              string     `db:"id"`
	Name            string     `db:"name"`
	Email           string     `db:"email"`
	PasswordHash    string     `db:"password_hash"`
	EmailVerifiedAt *time.Time `db:"email_verified_at"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
}
