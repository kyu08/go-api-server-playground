package usecase

import (
	"context"
	"database/sql"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/domain/service"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

type (
	CreateUserUsecase struct {
		db             *sql.DB
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

	if err := database.WithTransaction(ctx, u.db, func(queries *database.Queries) error {
		// NOTE: UserService内でTransactionを使うために必要なので注意
		u.userRepository.SetQueries(queries)
		userService := service.NewUserService(u.userRepository)
		return userService.CreateUser(ctx, newUser)
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &CreateUserOutput{
		ID: newUser.ID,
	}, nil
}

func NewCreateUserUsecase(db *sql.DB, userRepository *repository.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		db:             db,
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
