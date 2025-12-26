package database

import (
	"time"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

type User struct {
	ID         string
	ScreenName string
	UserName   string
	Bio        string
	IsPrivate  bool
	CreatedAt  time.Time
}

func (u *User) ToUser() *user.User {
	return &user.User{
		ID:         id.ID(u.ID),
		ScreenName: user.ScreenName(u.ScreenName),
		UserName:   user.UserName(u.UserName),
		Bio:        user.Bio(u.Bio),
		IsPrivate:  u.IsPrivate,
		CreatedAt:  u.CreatedAt,
	}
}

func UserFromRow(row *spanner.Row) (*User, error) {
	var u User
	if err := row.Columns(&u.ID, &u.ScreenName, &u.UserName, &u.Bio, &u.IsPrivate, &u.CreatedAt); err != nil {
		return nil, errors.WithStack(err)
	}

	return &u, nil
}
