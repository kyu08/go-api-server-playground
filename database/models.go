// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Author struct {
	ID    int64
	Name2 string
	Bio   sql.NullString
}
