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

func (UserRepository) FindByScreenName(ctx context.Context, db *sql.DB, screenName string) (*user.User, error) {
	queries := database.New(db)

	u, err := queries.FindUserByScreenName(ctx, screenName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrFindUserByScreenNameUserNotFound // TODO: repository層のエラーに変換する
		}

		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return &user.User{
		ID:         u.ID,
		ScreenName: user.ScreenName(u.ScreenName),
		UserName:   u.UserName,
		Bio:        u.Bio,
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	}, nil
}
