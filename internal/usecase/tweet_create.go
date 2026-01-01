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
		ID domain.ID[tweet.Tweet]
	}
)

var (
	ErrTweetCreateAuthorIDRequired = apperrors.NewPreconditionError("author_id is required")
	ErrTweetCreateBodyRequired     = apperrors.NewPreconditionError("body is required")
)

func (u TweetCreateUsecase) Run(ctx context.Context, input *TweetCreateInput) (*TweetCreateOutput, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	userID, err := domain.NewFromString[user.User](input.AuthorID)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.FindByID(ctx, u.client.Single(), userID)
	if err != nil {
		return nil, err
	}

	newTweet, err := tweet.NewTweet(user.ID, input.Body)
	if err != nil {
		return nil, err
	}

	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, rwtx *spanner.ReadWriteTransaction) error {
		return u.tweetRepository.Create(ctx, rwtx, newTweet)
	}); err != nil {
		// TODO: ここのエラー変換ロジックはいずれ共通化することになりそう。(どこの層の責務かもちょっと考えたほうがよさそう)
		if apperrors.IsPrecondition(err) || apperrors.IsNotFound(err) {
			return nil, apperrors.WithStack(err)
		}

		return nil, apperrors.NewInternalError(err)
	}

	return &TweetCreateOutput{
		ID: newTweet.ID,
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

func NewTweetCreateInput(screenName, userName, bio string) *TweetCreateInput {
	return &TweetCreateInput{
		AuthorID: screenName,
		Body:     userName,
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
