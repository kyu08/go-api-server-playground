package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, db *sql.DB, user *User) error
	FindByScreenName(ctx context.Context, db *sql.DB, screenName string) (*User, error)
}
