package database

import (
	"context"
	"fmt"

	"github.com/kyu08/go-api-server-playground/internal/domain/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

func (q *Queries) FindUserByScreenName_(ctx context.Context, screenName string) (*user.User, error) {
	u, err := q.FindUserByScreenName(ctx, screenName)
	if err != nil {
		return nil, fmt.Errorf("q.FindUserByScreenName: %w", err)
	}

	return &user.User{
		ID:         id.ID(u.ID),
		ScreenName: user.ScreenName(u.ScreenName),
		UserName:   user.UserName(u.UserName),
		Bio:        user.Bio(u.Bio),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	}, nil
}
