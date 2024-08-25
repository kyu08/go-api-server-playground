package errors

import (
	stderrors "errors"

	goerrors "github.com/go-errors/errors"
)

type (
	ErrorType    int
	TwitterError struct {
		Type    ErrorType
		Message string
	}
)

const (
	// Internal は内部エラー。DBとの接続が切れた場合などが発生した場合に使用する。
	Internal ErrorType = iota
	// Precondition は引数が不正なエラー。リクエストの引数が不正な場合に使用する。(バリデーションに違反する引数や存在しないユーザーへのアクションなど)
	Precondition
)

func (e TwitterError) Error() string {
	return e.Message
}

func NewInternalError(err error) error {
	return &TwitterError{
		Type:    Internal,
		Message: err.Error(),
	}
}

func NewPreconditionError(message string) error {
	return &TwitterError{
		Type:    Precondition,
		Message: message,
	}
}

func (e TwitterError) isPreconditionError() bool {
	switch e.Type {
	case Precondition:
		return true
	case Internal:
		return false
	default:
		return false
	}
}

func IsPreconditionError(err error) bool {
	var terror *TwitterError
	if stderrors.As(err, &terror) {
		return terror.isPreconditionError()
	}
	return false
}

func WithStack(err error) error {
	return goerrors.Wrap(err, 0)
}
