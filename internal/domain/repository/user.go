package repository

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, tx domain.ReadWriteDB, u *user.User) error
	FindByScreenName(ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName) (*user.User, error)
}
