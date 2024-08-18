package handler

import (
	"context"
	"database/sql"
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

		return nil, err
	}

	log.Printf("GetUserByScreenName: %#v", u)

	return &api.GetUserByScreenNameResponse{
		Id:         u.ID,
		ScreenName: u.ScreenName,
		Name:       u.Name,
		Bio:        u.Bio,
	}, nil
}

// TODO: パッケージ構成をいい感じにする
func testSQL(ctx context.Context, db *sql.DB, s string) (*database.User, error) {
	// TODO: Newの中でやっていることを理解する
	queries := database.New(db)
	u, err := queries.GetUserByScreenName(ctx, s)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
