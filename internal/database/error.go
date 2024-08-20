package database

import (
	"database/sql"
	"errors"
)

// TODO:add UT
func IsNotFoundError(err error, entity string) bool {
	return errors.Is(err, NotFoundError{Entity: entity})
}

func IsNotFound(err error) bool {
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

func (e NotFoundError) Is(err error, target error) bool {
	var err_ *NotFoundError
	var target_ *NotFoundError
	if errors.As(err, err_) && errors.As(target, target_) {
		return err_.Entity == e.Entity
	}

	return false
}
