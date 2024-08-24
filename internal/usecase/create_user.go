package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kyu08/go-api-server-playground/internal/database"
	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/domain/id"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
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
	ErrCreateUserScreenNameRequired    = errors.New("screen name is required")
	ErrCreateUserUserNameRequired      = errors.New("user name is required")
	ErrCreateUserScreenNameAlreadyUsed = errors.New("the screen name specified is already used")
)

func (u CreateUserUsecase) Run(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	if err := input.validate(); err != nil {
		return nil, fmt.Errorf("input.validate: %w", err)
	}

	newUser, err := user.New(input.ScreenName, input.UserName, input.Bio)
	if err != nil {
		return nil, fmt.Errorf("user.NewUser: %w", err)
	}

	if err := database.WithTransaction(ctx, u.db, func(queries *database.Queries) error {
		isUnique, err := isUniqueScreenName(ctx, u.userRepository, queries, newUser.ScreenName)
		if err != nil {
			return fmt.Errorf("isUniqueScreenName: %w", err)
		}
		if !isUnique {
			return ErrCreateUserScreenNameAlreadyUsed
		}

		if err := u.userRepository.Create(ctx, queries, newUser); err != nil {
			return fmt.Errorf("u.userRepository.Create: %w", err)
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("database.WithTransaction: %w", err)
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
		return ErrCreateUserScreenNameRequired
	}
	if i.UserName == "" {
		return ErrCreateUserUserNameRequired
	}

	return nil
}
