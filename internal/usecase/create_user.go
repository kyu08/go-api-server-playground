package usecase

import (
	"context"
	"database/sql"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/domain/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
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
		isUnique, err := userHelper.isUniqueScreenName(ctx, u.userRepository, queries, newUser.ScreenName)
		if err != nil {
			return errors.WithStack(err)
		}
		if !isUnique {
			return errors.WithStack(ErrCreateUserScreenNameAlreadyUsed)
		}

		if err := u.userRepository.Create(ctx, queries, newUser); err != nil {
			return errors.WithStack(err)
		}

		return nil
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
