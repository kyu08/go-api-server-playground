package user

import "errors"

type Bio string

var (
	ErrBioRequired = errors.New("bio is required")
	ErrBioTooLong  = errors.New("bio is too long")
)

func NewUserBio(source string) (Bio, error) {
	s := Bio(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

func (s Bio) validate() error {
	if len(s) == 0 {
		return ErrBioRequired
	}

	const bioMaxLength = 160
	if bioMaxLength < len(s) {
		return ErrBioTooLong
	}

	return nil
}

func (s Bio) String() string {
	return string(s)
}
