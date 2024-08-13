package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	pb "github.com/kyu08/go-api-server-playground/pkg/grpc" // TODO: rename
	oursql "github.com/kyu08/go-api-server-playground/sql"  // TODO: rename
)

func (s *TwitterServer) CountAuthors(ctx context.Context, _ *pb.CountAuthorsRequest) (*pb.CountAuthorsResponse, error) {
	log.Printf("Received: %v", "CountAuthors")

	a, err := testSQL(ctx, s.db)
	if err != nil {
		log.Printf("err: %s", err)

		return nil, err
	}

	log.Printf("CountAuthors: %#v", a)

	return &pb.CountAuthorsResponse{
		Count: int64(len(a)),
	}, nil
}

func testSQL(ctx context.Context, db *sql.DB) ([]oursql.Author, error) {
	queries := oursql.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.ListAuthors: %w", err)
	}

	return authors, nil
}
