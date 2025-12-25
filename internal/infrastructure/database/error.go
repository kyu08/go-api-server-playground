package database

import (
	"errors"

	"google.golang.org/api/iterator"
)

func IsNotFoundFromDB(err error) bool {
	return errors.Is(err, iterator.Done)
}
