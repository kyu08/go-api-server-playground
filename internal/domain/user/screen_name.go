package user

import "errors"

type ScreenName string

var (
	ErrScreenNameRequired = errors.New("screen_name is required")
	ErrScreenNameTooLong  = errors.New("screen_name is too long")
)

func NewUserScreenName(source string) (ScreenName, error) {
	s := ScreenName(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

func (s ScreenName) validate() error {
	if len(s) == 0 {
		return ErrScreenNameRequired
	}

	const screenNameMaxLength = 20
	if screenNameMaxLength < len(s) {
		return ErrScreenNameTooLong
	}

	return nil
}

func (s ScreenName) String() string {
	return string(s)
}
