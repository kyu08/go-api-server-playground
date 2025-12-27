package service

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/domain/repository"
)

var ErrCreateUserScreenNameAlreadyUsed = apperrors.NewPreconditionError("the screen name specified is already used")

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// TODO: add UT
func (s UserService) CreateUser(ctx context.Context, tx *spanner.ReadWriteTransaction, user *user.User) error {
	isExisting, err := s.IsExistingScreenName(ctx, tx, user.ScreenName)
	if err != nil {
		return apperrors.WithStack(err)
	}

	if isExisting {
		return apperrors.WithStack(ErrCreateUserScreenNameAlreadyUsed)
	}

	if err := s.userRepository.Create(ctx, tx, user); err != nil {
		return apperrors.WithStack(err)
	}

	return nil
}

// TODO: add UT
func (s UserService) IsExistingScreenName(ctx context.Context, tx *spanner.ReadWriteTransaction, screenName user.ScreenName) (bool, error) {
	user, err := s.userRepository.FindByScreenName(ctx, tx, screenName)
	if err != nil {
		if apperrors.IsNotFound(err) {
			return false, nil
		}

		return false, apperrors.WithStack(err)
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}
