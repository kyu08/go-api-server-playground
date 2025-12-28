package repository

import (
	"errors"
)

func IsNotFound(err error) bool {
	var ye *yoError
	if errors.As(err, &ye) {
		return ye.NotFound()
	}
	return false
}
