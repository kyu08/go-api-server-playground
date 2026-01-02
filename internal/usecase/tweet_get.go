package usecase

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/tweet"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type (
	TweetGetUsecase struct {
		client          *spanner.Client
		tweetRepository tweet.TweetRepository
		userRepository  user.UserRepository
	}
	TweetGetInput struct {
		TweetID string
	}
	TweetGetOutput struct {
		TweetID domain.ID[tweet.Tweet]
		// TODO: この辺をBody型にすべきかどうか迷う。
		// 結果によってはIDもstringにすべきかもしれない。
		Body              string
		AuthorId          domain.ID[user.User]
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

	// var newTweetID domain.ID[tweet.Tweet]
	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, rwtx *spanner.ReadWriteTransaction) error {
		panic("implement me")
	}); err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &TweetGetOutput{
		// TODO: fill this struct
		TweetID:           tweetID,
		Body:              "",
		AuthorId:          domain.ID[user.User]{},
		AuthorScreenName:  "",
		AuthorDisplayName: "",
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
	}, nil
}

func NewTweetGetUsecase(
	client *spanner.Client,
	tweetRepository tweet.TweetRepository,
	userRepository user.UserRepository,
) *TweetGetUsecase {
	return &TweetGetUsecase{
		client: client,
		// TODO: これをtweetUserRepositoryにする？tweetDetailRepositoryとかにする？
		tweetRepository: tweetRepository,
		userRepository:  userRepository,
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
