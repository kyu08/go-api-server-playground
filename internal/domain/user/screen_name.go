package user

import "errors"

type userScreenName string

var (
	ErrScreenNameEmpty             = errors.New("screen_name is empty")
	ErrScreenNameMoreThanMaxLength = errors.New("screen_name is more than max length")
)

func NewUserScreenName(screenName string) (userScreenName, error) {
	s := userScreenName(screenName)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

func (s userScreenName) validate() error {
	if len(s) == 0 {
		return ErrScreenNameEmpty
	}

	const screenNameMaxLength = 20
	if screenNameMaxLength < len(s) {
		return ErrScreenNameMoreThanMaxLength
	}

	return nil
}
