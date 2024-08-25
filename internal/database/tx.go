package database

import (
	"context"
	"database/sql"
	"fmt"
)

func WithTransaction(ctx context.Context, db *sql.DB, fn func(q *Queries) error) error {
	tx, err := db.BeginTx(ctx, nil) // TODO: トランザクション分離レベルをどうするか検討する
	if err != nil {
		return fmt.Errorf("db.BeginTx: %w", err)
	}

	defer func() { _ = tx.Rollback }() // コミットされた場合はRollbackされないのでdeferで問題ない

	if err := fn(New(tx)); err != nil {
		return fmt.Errorf("fn: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}
