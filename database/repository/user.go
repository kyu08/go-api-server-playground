package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

var ErrFindUserByScreenNameUserNotFound = errors.New("user not found")

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// TODO: internalディレクトリに移動する

func (r UserRepository) Create(ctx context.Context, queries *database.Queries, u *user.User) error {
	p := database.CreateUserParams{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	}
	if _, err := queries.CreateUser(ctx, p); err != nil {
		return fmt.Errorf("queries.CreateUser: %w", err)
	}

	return nil
}

func (UserRepository) FindByScreenName(
	ctx context.Context,
	queries *database.Queries,
	screenName user.ScreenName,
) (*user.User, error) {
	u, err := queries.FindUserByScreenName(ctx, string(screenName))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrFindUserByScreenNameUserNotFound // TODO: repository層のエラーに変換する
		}

		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return u.ToUser(), nil
}
