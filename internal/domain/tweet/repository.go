package tweet

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/domain"
)

type TweetRepository interface {
	Create(ctx context.Context, rwtx domain.ReadWriteDB, t *Tweet) error
}
