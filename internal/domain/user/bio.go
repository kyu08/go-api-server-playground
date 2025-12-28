package user

import "github.com/kyu08/go-api-server-playground/internal/apperrors"

type Bio string

var (
	ErrBioRequired = apperrors.NewPreconditionError("bio is required")
	ErrBioTooLong  = apperrors.NewPreconditionError("bio is too long")
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
		return apperrors.WithStack(ErrBioRequired)
	}

	const bioMaxLength = 160
	if bioMaxLength < len(s) {
		return apperrors.WithStack(ErrBioTooLong)
	}

	return nil
}

func (s Bio) String() string {
	return string(s)
}
