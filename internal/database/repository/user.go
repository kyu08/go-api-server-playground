package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

var ErrFindUserByScreenNameUserNotFound = errors.New("user not found")

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

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
		if database.IsNotFound(err) {
			return nil, database.NewNotFoundError("user")
		}

		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return u.ToUser(), nil
}
