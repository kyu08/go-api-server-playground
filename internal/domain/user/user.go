package user

import (
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain"
)

// User はユーザーを表すエンティティ。
// 基本的にフィールドはprivateにしてドメインロジックの流出を防ぐ。
// IDやCreatedAtなどの更新されないフィールドに関してはpublicにしている。
// これらのフィールドは基本的に変更されないのでpublicにするデメリット(安全性の低下)が少ないと判断したため。
type User struct {
	ID         domain.ID[User]
	screenName ScreenName
	userName   UserName
	bio        Bio
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
		screenName: s,
		userName:   u,
		bio:        b,
		CreatedAt:  createdAt,
	}
	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) ScreenName() ScreenName {
	return u.screenName
}

func (u *User) UserName() UserName {
	return u.userName
}

func (u *User) Bio() Bio {
	return u.bio
}

func (u *User) validate() error {
	// NOTE: 値オブジェクトのバリデーションでは収まらないエンティティとして別途必要なバリデーションが生まれたらここに書く想定
	return nil
}
