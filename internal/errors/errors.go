package errors

import (
	stderrors "errors"
	"fmt"

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
	// internal は内部エラー。DBとの接続が切れた場合などクライアント側が対処できないエラーを表す。エラー内容はクライアントには通知せず、ログにのみ出力する。
	internal ErrorType = iota
	// precondition は引数が不正なエラー。リクエストの引数が不正な場合に使用する。(バリデーションに違反する引数や存在しないユーザーへのアクションなど)
	precondition
	// notFound はリソースが見つからないエラー。リソースが存在しない場合に使用する。
	notFound
)

func (e TwitterError) Error() string {
	if !e.isPreconditionError() {
		return "Internal Error: " + e.Message
	}

	return e.Message
}

func NewInternalError(err error) error {
	return &TwitterError{
		Type:    internal,
		Message: err.Error(),
	}
}

func NewPreconditionError(message string) error {
	return &TwitterError{
		Type:    precondition,
		Message: message,
	}
}

func NewNotFoundError(entity string) error {
	return &TwitterError{
		Type:    notFound,
		Message: entity,
	}
}

func (e TwitterError) isPreconditionError() bool {
	switch e.Type {
	case precondition:
		return true
	case internal, notFound:
		return false
	default:
		return false
	}
}

func IsPrecondition(err error) bool {
	var terror *TwitterError
	if stderrors.As(err, &terror) {
		return terror.isPreconditionError()
	}

	return false
}

func WithStack(err error) error {
	return goerrors.Wrap(err, 0)
}

func GetStackTrace(err error) string {
	var gerr *goerrors.Error
	if stderrors.As(err, &gerr) {
		return fmt.Sprint(gerr.StackFrames())
	}

	return err.Error()
}

func IsNotFound(err error) bool {
	var terror *TwitterError
	if stderrors.As(err, &terror) {
		return terror.Type == notFound
	}

	return false
}
