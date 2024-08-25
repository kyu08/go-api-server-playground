package errors

import (
	stdErrors "errors"

	goerrors "github.com/go-errors/errors"
)

// ErrorType
type ErrorType int

const (
	// Unknown は不明なエラー。誤ってゼロ値で初期化された場合はこのエラーになるため基本的には出現しないはず。
	Unknown ErrorType = iota
	// Internal は内部エラー。DBとの接続が切れた場合などが発生した場合に使用する。
	Internal
	// Precondition は引数が不正なエラー。リクエストの引数が不正な場合に使用する。(バリデーションに違反する引数や存在しないユーザーへのアクションなど)
	Precondition

	InternalErrorMessage = "internal server error"
)

type TwitterError struct {
	Type    ErrorType
	Message string
}

func (e TwitterError) Error() string {
	return e.Message
}

func NewInternalError() error {
	return &TwitterError{
		Type:    Internal,
		Message: InternalErrorMessage,
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
	case Internal, Unknown:
		return false
	default:
		return false
	}
}

func IsPreconditionError(err error) bool {
	var terror *TwitterError
	if stdErrors.As(err, &terror) {
		return terror.isPreconditionError()
	}
	return false
}

func WithStack(err error) error {
	return goerrors.Wrap(err, 0)
}
