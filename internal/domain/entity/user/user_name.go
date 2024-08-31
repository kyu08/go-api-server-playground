package user

import "github.com/kyu08/go-api-server-playground/internal/errors"

type UserName string

var (
	ErrUserNameRequired = errors.NewPreconditionError("user_name is required")
	ErrUserNameTooLong  = errors.NewPreconditionError("user_name is too long")
)

func NewUserUserName(source string) (UserName, error) {
	s := UserName(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

func (s UserName) validate() error {
	if len(s) == 0 {
		return errors.WithStack(ErrUserNameRequired)
	}

	const userNameMaxLength = 20
	if userNameMaxLength < len(s) {
		return errors.WithStack(ErrUserNameTooLong)
	}

	return nil
}

func (s UserName) String() string {
	return string(s)
}
