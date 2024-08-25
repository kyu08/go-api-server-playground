package user

import "errors"

type UserName string

var (
	ErrUserNameRequired = errors.New("user_name is required")
	ErrUserNameTooLong  = errors.New("user_name is too long")
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
		return ErrUserNameRequired
	}

	const userNameMaxLength = 20
	if userNameMaxLength < len(s) {
		return ErrUserNameTooLong
	}

	return nil
}

func (s UserName) String() string {
	return string(s)
}
