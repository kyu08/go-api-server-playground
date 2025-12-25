package database

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

func WithTransaction(ctx context.Context, client *spanner.Client, fn func(txn *spanner.ReadWriteTransaction) error) error {
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		return fn(txn)
	})
	if err != nil {
		return errors.WithStack(errors.NewInternalError(err))
	}

	return nil
}
