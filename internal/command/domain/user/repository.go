package user

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/command/domain"
	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure"
)

type UserRepository interface {
	Create(ctx context.Context, rwtx infrastructure.ReadWriteDB, u *User) error
	FindByID(ctx context.Context, rtx infrastructure.ReadOnlyDB, userID domain.ID[User]) (*User, error)
	FindByScreenName(ctx context.Context, rtx infrastructure.ReadOnlyDB, screenName ScreenName) (*User, error)
}
