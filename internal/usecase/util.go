package usecase

import "github.com/kyu08/go-api-server-playground/internal/apperrors"

func handleError(err error) error {
	if apperrors.IsPrecondition(err) || apperrors.IsNotFound(err) {
		return err
	}

	return apperrors.NewInternalError(err)
}
