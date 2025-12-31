package user

import (
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type ScreenName domain.StringVO[ScreenName]

var (
	ErrScreenNameRequired = apperrors.NewPreconditionError("screen_name is required")
	ErrScreenNameTooLong  = apperrors.NewPreconditionError("screen_name is too long")
)

func NewUserScreenName(source string) (ScreenName, error) {
	s := domain.NewStringVOFromString[ScreenName](source)

	if err := s.validate(); err != nil {
		return ScreenName{}, apperrors.WithStack(err)
	}

	return s, nil
}

func (s ScreenName) validate() error {
	if len(s) == 0 {
		return apperrors.WithStack(ErrScreenNameRequired)
	}

	const screenNameMaxLength = 20
	if screenNameMaxLength < len(s) {
		return apperrors.WithStack(ErrScreenNameTooLong)
	}

	return nil
}

func (s ScreenName) String() string {
	return string(s)
}
