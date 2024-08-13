package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) CountAuthors(
	ctx context.Context,
	_ *api.CountAuthorsRequest,
) (*api.CountAuthorsResponse, error) {
	log.Printf("Received: %v", "CountAuthors")

	a, err := testSQL(ctx, s.db)
	if err != nil {
		log.Printf("err: %s", err)

		return nil, err
	}

	log.Printf("CountAuthors: %#v", a)

	return &api.CountAuthorsResponse{
		Count: int64(len(a)),
	}, nil
}

func testSQL(ctx context.Context, db *sql.DB) ([]database.Author, error) {
	queries := database.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.ListAuthors: %w", err)
	}

	return authors, nil
}
