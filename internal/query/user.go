package query

import (
	"context"
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type UserQuery interface {
	FindByID(ctx context.Context, rtx domain.ReadOnlyDB, userID domain.ID[user.User]) (*User, error)
	FindByScreenName(ctx context.Context, rtx domain.ReadOnlyDB, screenName user.ScreenName) (*User, error)
}

type (
	User struct {
		ID         string
		ScreenName string
		UserName   string
		Bio        string
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
