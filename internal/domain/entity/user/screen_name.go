package user

import "github.com/kyu08/go-api-server-playground/internal/errors"

type ScreenName string

var (
	ErrScreenNameRequired = errors.NewPreconditionError("screen_name is required")
	ErrScreenNameTooLong  = errors.NewPreconditionError("screen_name is too long")
)

func NewUserScreenName(source string) (ScreenName, error) {
	s := ScreenName(source)

	if err := s.validate(); err != nil {
		return "", errors.WithStack(err)
	}

	return s, nil
}

func (s ScreenName) validate() error {
	if len(s) == 0 {
		return errors.WithStack(ErrScreenNameRequired)
	}

	const screenNameMaxLength = 20
	if screenNameMaxLength < len(s) {
		return errors.WithStack(ErrScreenNameTooLong)
	}

	return nil
}

func (s ScreenName) String() string {
	return string(s)
}
