package usecase

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/domain/service"
	"github.com/kyu08/go-api-server-playground/internal/errors"
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
	ErrCreateUserScreenNameRequired    = errors.NewPreconditionError("screen name is required")
	ErrCreateUserUserNameRequired      = errors.NewPreconditionError("user name is required")
	ErrCreateUserScreenNameAlreadyUsed = errors.NewPreconditionError("the screen name specified is already used")
)

func (u CreateUserUsecase) Run(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	if err := input.validate(); err != nil {
		return nil, errors.WithStack(err)
	}

	newUser, err := user.New(input.ScreenName, input.UserName, input.Bio)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if _, err := u.client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// NOTE: UserService内でTransactionを使うために必要なので注意
		u.userRepository.SetTransaction(txn)
		userService := service.NewUserService(u.userRepository)
		return userService.CreateUser(ctx, newUser)
	}); err != nil {
		if errors.IsPrecondition(err) || errors.IsNotFound(err) {
			return nil, err
		}
		return nil, errors.WithStack(errors.NewInternalError(err))
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
		return errors.WithStack(ErrCreateUserScreenNameRequired)
	}
	if i.UserName == "" {
		return errors.WithStack(ErrCreateUserUserNameRequired)
	}

	return nil
}
