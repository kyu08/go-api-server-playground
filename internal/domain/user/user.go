package user

import (
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain/id"
)

// TODO: validateメソッド、UTをかく
// TODO: CQRS的に実装した方がよさそうか検討してみる。課題はREADではvalidationが必要ないのにコードの構造上それが
// 強制されそうで可読性が落ちたり無駄な計算が発生してしまう。READには必要ないフィールドだけどvalidationのために
// 取得する、というようなことも起きてしまいそう。（これは無駄に感じる）
// -> これはdatabase層のtoUserメソッドを使うことで解決できた。
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
		ID:         id.New(),
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

// TODO: UT
func (u *User) validate() error {
	if err := u.ScreenName.validate(); err != nil {
		return err
	}

	if err := u.UserName.validate(); err != nil {
		return err
	}

	if err := u.Bio.validate(); err != nil {
		return err
	}

	return nil
}
