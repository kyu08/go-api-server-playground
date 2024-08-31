package database

import (
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
)

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
