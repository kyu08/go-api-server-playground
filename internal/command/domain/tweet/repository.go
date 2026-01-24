package tweet

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure"
)

type TweetRepository interface {
	Create(ctx context.Context, rwtx infrastructure.ReadWriteDB, t *Tweet) error
}
