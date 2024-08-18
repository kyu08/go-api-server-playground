package user

import "errors"

type screenName string

var (
	ErrScreenNameEmpty             = errors.New("screen_name is empty")
	ErrScreenNameMoreThanMaxLength = errors.New("screen_name is more than max length")
)

func NewUserScreenName(source string) (screenName, error) {
	s := screenName(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

func (s screenName) validate() error {
	if len(s) == 0 {
		return ErrScreenNameEmpty
	}

	const screenNameMaxLength = 20
	if screenNameMaxLength < len(s) {
		return ErrScreenNameMoreThanMaxLength
	}

	return nil
}
