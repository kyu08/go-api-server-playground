package user

import (
	"context"
)

type UserRepository interface {
	FindUserByScreenName_(ctx context.Context, screenName string) (*User, error)
}
