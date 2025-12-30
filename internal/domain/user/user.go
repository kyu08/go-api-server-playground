package user

import (
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type User struct {
	ID         domain.ID[User]
	ScreenName ScreenName
	UserName   UserName
	Bio        Bio
	CreatedAt  time.Time
}

// NewUser はユーザー作成時に使用するコンストラクタ。
func NewUser(screenName, userName, bio string) (*User, error) {
	return newUser(domain.NewID[User](), screenName, userName, bio, time.Now())
}

// NewFromDTO は主にDTOからエンティティを生成する際に使用されることを想定し、IDを外から受け取るようにしている。
// これがドメイン知識かと言われると微妙だが、レイヤー間の依存関係の都合でこうするのが良いと判断した。
func NewFromDTO(idString string, screenName, userName, bio string, createdAt time.Time) (*User, error) {
	id, err := domain.NewFromString[User](idString)
	if err != nil {
		return nil, err
	}

	return newUser(id, screenName, userName, bio, createdAt)
}

func newUser(id domain.ID[User], screenName, userName, bio string, createdAt time.Time) (*User, error) {
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
		ID:         id,
		ScreenName: s,
		UserName:   u,
		Bio:        b,
		CreatedAt:  createdAt,
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
