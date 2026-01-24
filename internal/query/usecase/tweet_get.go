package usecase

import (
	"context"
	"strconv"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/query"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
	"github.com/samber/lo"
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
		AuthorID          string
		AuthorScreenName  string
		AuthorDisplayName string
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}
)

var ErrTweetGetTweetIDRequired = apperrors.NewPreconditionError("tweet_id is required")
var ErrTweetGetInvalidTweetID = apperrors.NewPreconditionError("invalid UUID length: ")

// ID指定でtweet詳細を1件取得する
func (u TweetGetUsecase) Run(ctx context.Context, input *TweetGetInput) (*TweetGetOutput, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	// UUIDのバリデーション
	if err := validateUUID(input.TweetID); err != nil {
		return nil, err
	}

	res, err := u.tweetQuery.GetDetail(ctx, u.client.Single(), input.TweetID)
	if err != nil {
		return nil, err
	}

	return lo.ToPtr(TweetGetOutput(*res)), nil
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

func validateUUID(s string) error {
	// UUID形式の簡易バリデーション (長さチェック)
	if len(s) != 36 {
		return apperrors.WithStack(apperrors.NewPreconditionError("invalid UUID length: " + strconv.Itoa(len(s))))
	}
	return nil
}
