package commandimpl

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/tweet"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/model"
)

type TweetRepository struct{}

func NewTweetRepository() tweet.TweetRepository {
	return &TweetRepository{}
}

func (r TweetRepository) Create(ctx context.Context, rwtx domain.ReadWriteDB, u *tweet.Tweet) error {
	return r.apply(rwtx, []*spanner.Mutation{r.fromDomain(u).Insert(ctx)})
}

func (TweetRepository) apply(rwtx domain.ReadWriteDB, m []*spanner.Mutation) error {
	if err := rwtx.BufferWrite(m); err != nil {
		return apperrors.WithStack(apperrors.NewInternalError(err))
	}
	return nil
}

func (TweetRepository) fromDomain(u *tweet.Tweet) *model.Tweet {
	return &model.Tweet{
		ID:        u.ID.String(),
		AuthorID:  u.AuthorID.String(),
		Body:      u.Body().String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
