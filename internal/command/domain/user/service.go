package user

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
)

var ErrCreateUserScreenNameAlreadyUsed = apperrors.NewPreconditionError("the screen name specified is already used")

type CreateUserService struct {
	userRepository UserRepository
}

func NewCreateUserService(userRepository UserRepository) *CreateUserService {
	return &CreateUserService{
		userRepository: userRepository,
	}
}

// TODO: add UT
func (s CreateUserService) CreateUser(ctx context.Context, rwtx *spanner.ReadWriteTransaction, user *User) error {
	isExisting, err := s.IsExistingScreenName(ctx, rwtx, user.ScreenName())
	if err != nil {
		return apperrors.WithStack(err)
	}

	if isExisting {
		return apperrors.WithStack(ErrCreateUserScreenNameAlreadyUsed)
	}

	if err := s.userRepository.Create(ctx, rwtx, user); err != nil {
		return apperrors.WithStack(err)
	}

	return nil
}

// TODO: add UT
func (s CreateUserService) IsExistingScreenName(
	ctx context.Context,
	rwtx *spanner.ReadWriteTransaction,
	screenName ScreenName,
) (bool, error) {
	user, err := s.userRepository.FindByScreenName(ctx, rwtx, screenName)
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
