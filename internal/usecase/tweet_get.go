package usecase

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/tweet"
	"github.com/kyu08/go-api-server-playground/internal/query"
)

type (
	TweetGetUsecase struct {
		client     *spanner.Client
		tweetQuery query.TweetQuery
	}
	TweetGetInput struct {
		TweetID string
	}
	TweetGetOutput struct {
		TweetID           string
		Body              string
		AuthorId          string
		AuthorScreenName  string
		AuthorDisplayName string
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}
)

var ErrTweetGetTweetIDRequired = apperrors.NewPreconditionError("tweet_id is required")

// ID指定でtweet詳細を1件取得する
func (u TweetGetUsecase) Run(ctx context.Context, input *TweetGetInput) (*TweetGetOutput, error) {
	res, err := u.run(ctx, input)
	if err != nil {
		return nil, handleError(err)
	}

	return res, nil
}

func (u TweetGetUsecase) run(ctx context.Context, input *TweetGetInput) (*TweetGetOutput, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	tweetID, err := domain.NewFromString[tweet.Tweet](input.TweetID)
	if err != nil {
		return nil, err
	}

	res, err := u.tweetQuery.GetDetail(ctx, u.client.Single(), tweetID.String())
	if err != nil {
		return nil, err
	}

	return &TweetGetOutput{
		TweetID:           res.TweetID,
		Body:              res.Body,
		AuthorId:          res.AuthorID,
		AuthorScreenName:  res.AuthorScreenName,
		AuthorDisplayName: res.AuthorDisplayName,
		CreatedAt:         res.CreatedAt,
		UpdatedAt:         res.UpdatedAt,
	}, nil
}

func NewTweetGetUsecase(
	client *spanner.Client,
	tweetQuery query.TweetQuery,
) *TweetGetUsecase {
	return &TweetGetUsecase{
		client:     client,
		tweetQuery: tweetQuery,
	}
}

func NewTweetGetInput(tweetID string) *TweetGetInput {
	return &TweetGetInput{
		TweetID: tweetID,
	}
}

func (i TweetGetInput) validate() error {
	if i.TweetID == "" {
		return apperrors.WithStack(ErrTweetGetTweetIDRequired)
	}

	return nil
}
