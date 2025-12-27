package database

import (
	"errors"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func IsNotFoundFromDB(err error) bool {
	return errors.Is(err, spanner.ErrRowNotFound) || errors.Is(err, iterator.Done)
}
