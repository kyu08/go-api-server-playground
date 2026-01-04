package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/query"
)

type (
	FindUserByScreenNameUsecase struct {
		client    *spanner.Client
		userQuery query.UserQuery
	}
	FindUserByScreenNameInput struct {
		ScreenName string
	}
	FindUserByScreenNameOutput struct {
		ID         string
		ScreenName string
		UserName   string
		Bio        string
	}
)

var (
	ErrFindUserByScreenNameScreenNameRequired = apperrors.NewPreconditionError("screen name is required")
	ErrFindUserByScreenNameUserNotFound       = apperrors.NewNotFoundError("user")
)

// ScreeName指定でユーザーを1件取得する
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

	rtx := u.client.Single()
	foundUser, err := u.userQuery.FindByScreenName(ctx, rtx, screenName)
	if err != nil {
		if apperrors.IsNotFound(err) {
			return nil, apperrors.WithStack(ErrFindUserByScreenNameUserNotFound)
		}

		if apperrors.IsPrecondition(err) {
			return nil, apperrors.WithStack(err)
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
	userQuery query.UserQuery,
) *FindUserByScreenNameUsecase {
	return &FindUserByScreenNameUsecase{
		client:    client,
		userQuery: userQuery,
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
