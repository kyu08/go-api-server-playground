package user

import "errors"

type Bio string

var (
	ErrBioEmpty             = errors.New("bio is empty")
	ErrBioMoreThanMaxLength = errors.New("bio is more than max length")
)

func NewUserBio(source string) (Bio, error) {
	s := Bio(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

// TODO: UT
func (s Bio) validate() error {
	if len(s) == 0 {
		return ErrBioEmpty
	}

	const bioMaxLength = 160
	if bioMaxLength < len(s) {
		return ErrBioMoreThanMaxLength
	}

	return nil
}

func (s Bio) String() string {
	return string(s)
}
