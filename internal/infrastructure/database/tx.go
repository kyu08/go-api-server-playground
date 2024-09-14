package database

import (
	"context"
	"database/sql"

	"github.com/kyu08/go-api-server-playground/internal/errors"
)

func WithTransaction(ctx context.Context, db *sql.DB, fn func(q *Queries) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WithStack(errors.NewInternalError(err))
	}

	defer func() { _ = tx.Rollback }() // コミットされた場合はRollbackされないのでdeferで問題ない

	if err := fn(New(tx)); err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(errors.NewInternalError(err))
	}

	return nil
}
