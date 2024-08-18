package user

import "errors"

type UserName string

var (
	ErrUserNameEmpty             = errors.New("user_name is empty")
	ErrUserNameMoreThanMaxLength = errors.New("user_name is more than max length")
)

func NewUserUserName(source string) (UserName, error) {
	s := UserName(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

// TODO: UT
func (s UserName) validate() error {
	if len(s) == 0 {
		return ErrUserNameEmpty
	}

	const userNameMaxLength = 20
	if userNameMaxLength < len(s) {
		return ErrUserNameMoreThanMaxLength
	}

	return nil
}

func (s UserName) String() string {
	return string(s)
}
