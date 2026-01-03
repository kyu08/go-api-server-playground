package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/tweet"
	"github.com/kyu08/go-api-server-playground/internal/query"
)

// TODO: のちほどTweetQueryに統合する
type TweetRepository struct{}

func NewTweetRepository() tweet.TweetRepository {
	return &TweetRepository{}
}

type TweetQuery struct{}

func NewTweetQuery() tweet.TweetQuery {
	return &TweetQuery{}
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

func (TweetRepository) fromDomain(u *tweet.Tweet) *Tweet {
	return &Tweet{
		ID:        u.ID.String(),
		AuthorID:  u.AuthorID.String(),
		Body:      u.Body().String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (TweetQuery) GetDetail(ctx context.Context, rtx domain.ReadOnlyDB, tweetID string) ([]*query.TweetDetail, error) {
	queryStr := `
	SELECT
		TweetID,
		Body,
		AuthorID,
		AuthorScreenName,
		AuthorDisplayName,
		CreatedAt,
		UpdatedAt
	FROM 
		Tweet as t
	INNER JOIN 
		User as u
		ON u.ID = t.AuthorID
	WHERE
		t.ID = @tweetID
	`

	statement := spanner.NewStatement(queryStr)
	statement.Params["tweetID"] = tweetID

	iter := rtx.Query(ctx, statement)
	defer iter.Stop()

	result, err := toStruct[query.TweetDetail](iter)
	if err != nil {
		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}

	return result, nil
}

// func (TweetRepository) toDomain(dto *Tweet) (*tweet.Tweet, error) {
// 	u, err := tweet.NewFromDTO(
// 		dto.ID,
// 		dto.AuthorID,
// 		dto.Body,
// 		dto.CreatedAt,
// 		dto.UpdatedAt,
// 	)
// 	if err != nil {
// 		return nil, apperrors.WithStack(err)
// 	}
// 	return u, nil
// }
