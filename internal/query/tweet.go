package query

import (
	"context"
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type TweetQuery interface {
	GetDetail(ctx context.Context, rtx domain.ReadOnlyDB, tweetID string) (*TweetDetail, error)
}

type (
	TweetDetail struct {
		TweetID           string
		Body              string
		AuthorID          string
		AuthorScreenName  string
		AuthorDisplayName string
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}
)
