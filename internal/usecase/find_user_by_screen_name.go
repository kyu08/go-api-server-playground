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
	FindUserByScreenNameUsecase struct {
		db             *sql.DB
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

var ErrFindUserByScreenNameScreenNameRequired = errors.NewPreconditionError("screen name is required")

func (u FindUserByScreenNameUsecase) Run(
	ctx context.Context,
	input *FindUserByScreenNameInput,
) (*FindUserByScreenNameOutput, error) {
	if err := input.validate(); err != nil {
		return nil, errors.WithStack(err)
	}

	screenName, err := user.NewUserScreenName(input.ScreenName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queries := database.New(u.db)
	user, err := u.userRepository.FindByScreenName(ctx, queries, screenName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &FindUserByScreenNameOutput{
		ID:         user.ID,
		ScreenName: user.ScreenName,
		UserName:   user.UserName,
		Bio:        user.Bio,
	}, nil
}

func NewFindUserByScreenNameUsecase(
	db *sql.DB,
	userRepository *repository.UserRepository,
) *FindUserByScreenNameUsecase {
	return &FindUserByScreenNameUsecase{
		db:             db,
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
		return errors.WithStack(ErrFindUserByScreenNameScreenNameRequired)
	}

	return nil
}