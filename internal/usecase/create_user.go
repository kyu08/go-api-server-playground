package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type (
	CreateUserUsecase struct {
		client         *spanner.Client
		userRepository user.UserRepository
	}
	CreateUserInput struct {
		ScreenName string
		UserName   string
		Bio        string
	}
	CreateUserOutput struct {
		ID id.ID
	}
)

var (
	ErrCreateUserScreenNameRequired    = apperrors.NewPreconditionError("screen name is required")
	ErrCreateUserUserNameRequired      = apperrors.NewPreconditionError("user name is required")
	ErrCreateUserScreenNameAlreadyUsed = apperrors.NewPreconditionError("the screen name specified is already used")
)

func (u CreateUserUsecase) Run(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	if err := input.validate(); err != nil {
		return nil, err
	}

	newUser, err := user.New(input.ScreenName, input.UserName, input.Bio)
	if err != nil {
		return nil, err
	}

	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, rwtx *spanner.ReadWriteTransaction) error {
		userService := user.NewUserService(u.userRepository)
		return userService.CreateUser(ctx, rwtx, newUser)
	}); err != nil {
		// TODO: ここのエラー変換ロジックはいずれ共通化することになりそう。(どこの層の責務かもちょっと考えたほうがよさそう)
		if apperrors.IsPrecondition(err) || apperrors.IsNotFound(err) {
			return nil, apperrors.WithStack(err)
		}

		return nil, apperrors.NewInternalError(err)
	}

	return &CreateUserOutput{
		ID: newUser.ID,
	}, nil
}

func NewCreateUserUsecase(client *spanner.Client, userRepository user.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		client:         client,
		userRepository: userRepository,
	}
}

func NewCreateUserInput(screenName, userName, bio string) *CreateUserInput {
	return &CreateUserInput{
		ScreenName: screenName,
		UserName:   userName,
		Bio:        bio,
	}
}

func (i CreateUserInput) validate() error {
	if i.ScreenName == "" {
		return apperrors.WithStack(ErrCreateUserScreenNameRequired)
	}

	if i.UserName == "" {
		return apperrors.WithStack(ErrCreateUserUserNameRequired)
	}

	return nil
}
