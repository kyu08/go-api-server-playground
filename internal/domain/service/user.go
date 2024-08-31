package service

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain/entity/user"
	"github.com/kyu08/go-api-server-playground/internal/domain/repository"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

var ErrCreateUserScreenNameAlreadyUsed = errors.NewPreconditionError("the screen name specified is already used")

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// TODO: add UT
func (s UserService) CreateUser(ctx context.Context, user *user.User) error {
	isExisting, err := s.IsExistingScreenName(ctx, user.ScreenName)
	if err != nil {
		return errors.WithStack(err)
	}
	if isExisting {
		return errors.WithStack(ErrCreateUserScreenNameAlreadyUsed)
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// TODO: add UT
func (s UserService) IsExistingScreenName(ctx context.Context, screenName user.ScreenName) (bool, error) {
	user, err := s.userRepository.FindByScreenName(ctx, screenName)
	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, errors.WithStack(err)
	}

	if user != nil {
		return true, nil
	}
	return false, nil
}
