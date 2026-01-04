package queryimpl

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/query"
)

type TweetQuery struct{}

func NewTweetQuery() query.TweetQuery {
	return &TweetQuery{}
}

func (TweetQuery) GetDetail(ctx context.Context, rtx domain.ReadOnlyDB, tweetID string) (*query.TweetDetail, error) {
	queryStr := `
	SELECT
		t.ID AS TweetID,
		t.Body,
		t.AuthorID,
		u.ScreenName AS AuthorScreenName,
		u.UserName AS AuthorDisplayName,
		t.CreatedAt,
		t.UpdatedAt
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

	result, err := database.ToStruct[query.TweetDetail](iter)
	if err != nil {
		return nil, apperrors.WithStack(apperrors.NewInternalError(err))
	}
	if len(result) == 0 {
		return nil, apperrors.WithStack(database.NewNotFoundError[query.TweetDetail]())
	}

	return result[0], nil
}
