package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
)

type (
	FindUserByScreenNameUsecase struct {
		client         *spanner.Client
		userRepository *repository.UserRepository
	}
	FindUserByScreenNameInput struct {
		ScreenName string
	}
	FindUserByScreenNameOutput struct {
		ID         id.ID
		ScreenName user.ScreenName
		UserName   user.UserName
		Bio        user.Bio
	}
)

var (
	ErrFindUserByScreenNameScreenNameRequired = apperrors.NewPreconditionError("screen name is required")
	ErrFindUserByScreenNameUserNotFound       = apperrors.NewNotFoundError("user not found")
)

func (u FindUserByScreenNameUsecase) Run(
	ctx context.Context,
	input *FindUserByScreenNameInput,
) (*FindUserByScreenNameOutput, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	screenName, err := user.NewUserScreenName(input.ScreenName)
	if err != nil {
		return nil, err
	}

	var foundUser *user.User

	// Use ReadWriteTransaction to query (read-only operations also work within RW transaction)
	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		var findErr error
		foundUser, findErr = u.userRepository.FindByScreenName(ctx, tx, screenName)
		return findErr
	}); err != nil {
		if apperrors.IsNotFound(err) {
			return nil, apperrors.WithStack(ErrFindUserByScreenNameUserNotFound)
		}

		if apperrors.IsPrecondition(err) {
			return nil, err
		}

		return nil, apperrors.NewInternalError(err)
	}

	return &FindUserByScreenNameOutput{
		ID:         foundUser.ID,
		ScreenName: foundUser.ScreenName,
		UserName:   foundUser.UserName,
		Bio:        foundUser.Bio,
	}, nil
}

func NewFindUserByScreenNameUsecase(
	client *spanner.Client,
	userRepository *repository.UserRepository,
) *FindUserByScreenNameUsecase {
	return &FindUserByScreenNameUsecase{
		client:         client,
		userRepository: userRepository,
	}
}

func NewFindUserByScreenNameInput(screenName string) *FindUserByScreenNameInput {
	return &FindUserByScreenNameInput{
		ScreenName: screenName,
	}
}

func (i FindUserByScreenNameInput) validate() error {
	if i.ScreenName == "" {
		return apperrors.WithStack(ErrFindUserByScreenNameScreenNameRequired)
	}

	return nil
}
