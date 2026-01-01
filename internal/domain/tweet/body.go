package tweet

import (
	"unicode/utf8"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
)

type Body string

var (
	ErrBodyRequired = apperrors.NewPreconditionError("body is required")
	ErrBodyTooLong  = apperrors.NewPreconditionError("body is too long")
)

func NewBody(source string) (Body, error) {
	s := Body(source)

	if err := s.validate(); err != nil {
		return "", err
	}

	return s, nil
}

func (s Body) validate() error {
	if utf8.RuneCountInString(s.String()) == 0 {
		return apperrors.WithStack(ErrBodyRequired)
	}

	const bodyMaxLength = 140
	if bodyMaxLength < utf8.RuneCountInString(s.String()) {
		return apperrors.WithStack(ErrBodyTooLong)
	}

	return nil
}

func (s Body) String() string {
	return string(s)
}
