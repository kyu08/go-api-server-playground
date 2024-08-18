package user

import "time"

type User struct {
	// TODO: それぞれ値オブジェクトをつくる？
	ID         string
	ScreenName screenName
	UserName   string
	Bio        string
	IsPrivate  bool
	CreatedAt  time.Time
}
