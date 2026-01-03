package query

import "context"

type TweetQuery interface {
	GetDetail(ctx context.Context, tweetID string) (*TweetDetail, error)
}

type (
	TweetDetail struct {
		TweetID           string
		Body              string
		AuthorID          string
		AuthorScreenName  string
		AuthorDisplayName string
		CreatedAt         string
		UpdatedAt         string
	}
)
