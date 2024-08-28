package repository

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, u *user.User) error
	FindByScreenName(ctx context.Context, screenName user.ScreenName) (*user.User, error)
}
