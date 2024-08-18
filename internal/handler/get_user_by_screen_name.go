package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) GetUserByScreenName(
	ctx context.Context,
	_ *api.GetUserByScreenNameRequest,
) (*api.GetUserByScreenNameResponse, error) {
	log.Printf("Received: %v", "GetUserByScreenName")

	u, err := testSQL(ctx, s.db, "s_name")
	if err != nil {
		log.Printf("err: %s", err)

		return nil, fmt.Errorf("queries.GetUserByScreenName: %w", err)
	}

	log.Printf("GetUserByScreenName: %#v", u)

	return &api.GetUserByScreenNameResponse{
		Id:         u.ID,
		ScreenName: u.ScreenName,
		UserName:   u.UserName,
		Bio:        u.Bio,
	}, nil
}

// TODO: パッケージ構成をいい感じにする
func testSQL(ctx context.Context, db *sql.DB, s string) (*database.User, error) {
	// TODO: Newの中でやっていることを理解する
	queries := database.New(db)

	u, err := queries.GetUserByScreenName(ctx, s)
	if err != nil {
		return nil, fmt.Errorf("queries.GetUserByScreenName: %w", err)
	}

	return &u, nil
}
