package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, queries *database.Queries, user *user.User) error
	FindByScreenName(ctx context.Context, queries *database.Queries, screenName user.ScreenName) (*user.User, error)
}
