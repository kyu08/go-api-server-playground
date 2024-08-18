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

func (UserRepository) Create(ctx context.Context, db *sql.DB, u *user.User) error {
	queries := database.New(db)

	_, err := queries.CreateUser(ctx, database.CreateUserParams{
		ID:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateUser: %w", err)
	}

	return nil
}

func (UserRepository) FindByScreenName(ctx context.Context, db *sql.DB, screenName string) (*user.User, error) {
	queries := database.New(db)

	u, err := queries.FindUserByScreenName(ctx, screenName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrFindUserByScreenNameUserNotFound // TODO: repository層のエラーに変換する
		}

		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return u.ToUser(), nil
}
