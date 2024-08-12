package handler

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
	oursql "github.com/kyu08/go-api-server-playground/sql"
)

func (s *TwitterServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	log.Printf("Received: %v", "Health")

	testSQL(ctx, s.db)

	return &pb.HealthResponse{Message: "twitter"}, nil
}

func testSQL(ctx context.Context, db *sql.DB) {
	queries := oursql.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("authors: %v\n", authors)
}
