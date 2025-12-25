package database

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/config"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

func NewSpannerClient(ctx context.Context, config *config.Config) (*spanner.Client, error) {
	client, err := spanner.NewClient(ctx, config.DatabaseName())
	if err != nil {
		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	return client, nil
}
