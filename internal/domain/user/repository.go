package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	FindByScreenName(ctx context.Context, db *sql.DB, screenName string) (*User, error)
}
