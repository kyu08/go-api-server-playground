package user

import (
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain/id"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

type User struct {
	ID         id.ID
	ScreenName ScreenName
	UserName   UserName
	Bio        Bio
	IsPrivate  bool
	CreatedAt  time.Time
}

func New(screenName, userName, bio string) (*User, error) {
	s, err := NewUserScreenName(screenName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	u, err := NewUserUserName(userName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := NewUserBio(bio)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := &User{
		ID:         id.New(),
		ScreenName: s,
		UserName:   u,
		Bio:        b,
		IsPrivate:  false,
		CreatedAt:  time.Now(),
	}
	if err := user.validate(); err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

func (u *User) validate() error {
	// NOTE: 値オブジェクトのバリデーションでは収まらないエンティティとして別途必要なバリデーションが生まれたらここに書く想定
	return nil
}
