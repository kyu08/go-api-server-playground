package database

import (
	"errors"

	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/model"
)

func IsNotFound(err error) bool {
	var ye *model.YoError
	if errors.As(err, &ye) {
		return ye.NotFound()
	}
	return false
}
