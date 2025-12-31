package user

import (
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type ScreenName domain.StringVO[ScreenName]

var (
	ErrScreenNameRequired = apperrors.NewPreconditionError("screen_name is required")
	ErrScreenNameTooLong  = apperrors.NewPreconditionError("screen_name is too long")
)

func NewUserScreenName(source string) (ScreenName, error) {
	s := domain.NewStringVOFromString[ScreenName](source)

	if err := s.validate(); err != nil {
		return "", apperrors.WithStack(err)
	}

	return s, nil
}

func (s ScreenName) validate() error {
	if len(s) == 0 {
		return apperrors.WithStack(ErrScreenNameRequired)
	}

	const screenNameMaxLength = 20
	if screenNameMaxLength < len(s) {
		return apperrors.WithStack(ErrScreenNameTooLong)
	}

	return nil
}

func (s ScreenName) String() string {
	return string(s)
}

type (
	S1 domain.StringVO[S1]
	S2 domain.StringVO[S2]
	S3 string
	S4 string
)

func some() {
	s1 := domain.NewStringVOFromString[S1]("foo")
	s2 := S2(s1)

	f := func(s3 S3) {}
	s3 := S3(s1)
	s4 := S4(s3)

	f(s3)
}
