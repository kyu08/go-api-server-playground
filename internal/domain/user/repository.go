package user

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, rwtx domain.ReadWriteDB, u *User) error
	FindByScreenName(ctx context.Context, rtx domain.ReadOnlyDB, screenName ScreenName) (*User, error)
}
