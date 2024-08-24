package database

import (
	"database/sql"
	"errors"
)

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	target_ := new(NotFoundError)

	return errors.As(err, target_)
}

func IsNotFoundFromDB(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

type NotFoundError struct {
	Entity string
}

func NewNotFoundError(entity string) NotFoundError {
	return NotFoundError{Entity: entity}
}

func (e NotFoundError) Error() string {
	return e.Entity + " not found"
}
