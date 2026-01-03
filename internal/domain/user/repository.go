package user

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, rwtx domain.ReadWriteDB, u *User) error
	FindByID(ctx context.Context, rtx domain.ReadOnlyDB, userID domain.ID[User]) (*User, error)
}
