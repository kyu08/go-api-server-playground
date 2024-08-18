package user

import "context"

type UserRepository interface { // TODO: 定義側にinterfaceを定義するのはgo-wayではないな〜。どうするか検討する。
	FindUserByScreenName(ctx context.Context, screenName string) (*User, error)
}
