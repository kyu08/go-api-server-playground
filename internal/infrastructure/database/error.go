package database

import (
	"database/sql"
	"errors"
)

func IsNotFoundFromDB(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
