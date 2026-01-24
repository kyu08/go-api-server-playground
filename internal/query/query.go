package query

import (
	"context"
	"time"

	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure"
)

// UserQuery はユーザー検索用のインターフェース
// query側ではdomain.ID[T]を使わずstring IDを直接使用する
type UserQuery interface {
	FindByID(ctx context.Context, rtx infrastructure.ReadOnlyDB, userID string) (*User, error)
	FindByScreenName(ctx context.Context, rtx infrastructure.ReadOnlyDB, screenName string) (*User, error)
}

// User はユーザーの読み取りモデル
type User struct {
	ID         string
	ScreenName string
	UserName   string
	Bio        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TweetQuery はツイート検索用のインターフェース
type TweetQuery interface {
	GetDetail(ctx context.Context, rtx infrastructure.ReadOnlyDB, tweetID string) (*TweetDetail, error)
}

// TweetDetail はツイート詳細の読み取りモデル
type TweetDetail struct {
	TweetID           string
	Body              string
	AuthorID          string
	AuthorScreenName  string
	AuthorDisplayName string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
