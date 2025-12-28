package user

import (
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
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
	return NewWithID(id.New().String(), screenName, userName, bio)
}

// NewWithID は主にDTOからエンティティを生成する際に使用されることを想定し、IDを外から受け取るようにしている。
func NewWithID(idString string, screenName, userName, bio string) (*User, error) {
	s, err := NewUserScreenName(screenName)
	if err != nil {
		return nil, err
	}

	u, err := NewUserUserName(userName)
	if err != nil {
		return nil, err
	}

	b, err := NewUserBio(bio)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:         id.ID(idString),
		ScreenName: s,
		UserName:   u,
		Bio:        b,
		IsPrivate:  false,
		CreatedAt:  time.Now(),
	}
	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) validate() error {
	// NOTE: 値オブジェクトのバリデーションでは収まらないエンティティとして別途必要なバリデーションが生まれたらここに書く想定
	return nil
}
