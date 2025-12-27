package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/domain/service"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
)

type (
	CreateUserUsecase struct {
		client         *spanner.Client
		userRepository *repository.UserRepository
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
		return nil, apperrors.WithStack(err)
	}

	newUser, err := user.New(input.ScreenName, input.UserName, input.Bio)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}

	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
		userService := service.NewUserService(u.userRepository)
		return userService.CreateUser(ctx, tx, newUser)
	}); err != nil {
		if apperrors.IsPrecondition(err) || apperrors.IsNotFound(err) {
			return nil, apperrors.WithStack(err)
		}

		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return &CreateUserOutput{
		ID: newUser.ID,
	}, nil
}

func NewCreateUserUsecase(client *spanner.Client, userRepository *repository.UserRepository) *CreateUserUsecase {
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
