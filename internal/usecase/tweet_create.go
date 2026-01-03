package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/tweet"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type (
	TweetCreateUsecase struct {
		client          *spanner.Client
		tweetRepository tweet.TweetRepository
		userRepository  user.UserRepository
	}
	TweetCreateInput struct {
		AuthorID string
		Body     string
	}
	TweetCreateOutput struct {
		ID string
	}
)

var (
	ErrTweetCreateAuthorIDRequired = apperrors.NewPreconditionError("author_id is required")
	ErrTweetCreateBodyRequired     = apperrors.NewPreconditionError("body is required")
)

// tweetを作成する
func (u TweetCreateUsecase) Run(ctx context.Context, input *TweetCreateInput) (*TweetCreateOutput, error) {
	res, err := u.run(ctx, input)
	if err != nil {
		return nil, handleError(err)
	}

	return res, nil
}

func (u TweetCreateUsecase) run(ctx context.Context, input *TweetCreateInput) (*TweetCreateOutput, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	userID, err := domain.NewFromString[user.User](input.AuthorID)
	if err != nil {
		return nil, err
	}

	var newTweetID domain.ID[tweet.Tweet]
	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, rwtx *spanner.ReadWriteTransaction) error {
		user, err := u.userRepository.FindByID(ctx, rwtx, userID)
		if err != nil {
			return err
		}

		newTweet, err := tweet.NewTweet(user.ID, input.Body)
		if err != nil {
			return err
		}

		if err := u.tweetRepository.Create(ctx, rwtx, newTweet); err != nil {
			return err
		}

		newTweetID = newTweet.ID
		return nil
	}); err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &TweetCreateOutput{
		ID: newTweetID.String(),
	}, nil
}

func NewTweetCreateUsecase(
	client *spanner.Client,
	tweetRepository tweet.TweetRepository,
	userRepository user.UserRepository,
) *TweetCreateUsecase {
	return &TweetCreateUsecase{
		client:          client,
		tweetRepository: tweetRepository,
		userRepository:  userRepository,
	}
}

func NewTweetCreateInput(authorID string, body string) *TweetCreateInput {
	return &TweetCreateInput{
		AuthorID: authorID,
		Body:     body,
	}
}

func (i TweetCreateInput) validate() error {
	if i.AuthorID == "" {
		return apperrors.WithStack(ErrTweetCreateAuthorIDRequired)
	}

	if i.Body == "" {
		return apperrors.WithStack(ErrTweetCreateBodyRequired)
	}

	return nil
}
