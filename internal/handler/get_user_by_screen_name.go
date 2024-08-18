package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

var (
	ErrGetUserByScreenNameScreenNameEmpty             = errors.New("screen_name is empty")
	ErrGetUserByScreenNameScreenNameMoreThanMaxLength = errors.New("screen_name is more than max length")
	ErrGetUserByScreenNameUserNotFound                = errors.New("user not found")
)

func (s *TwitterServer) GetUserByScreenName(
	ctx context.Context,
	req *api.GetUserByScreenNameRequest,
) (*api.GetUserByScreenNameResponse, error) {
	log.Printf("Received: %v", "GetUserByScreenName")

	const screenNameMaxLength = 20

	if req.GetScreenName() == "" {
		return nil, ErrGetUserByScreenNameScreenNameEmpty
	}

	if screenNameMaxLength < len(req.GetScreenName()) {
		return nil, ErrGetUserByScreenNameScreenNameMoreThanMaxLength
	}

	// TODO: 値オブジェクトを使ってvalidate
	u, err := testSQL(ctx, s.db, req.GetScreenName())
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
	queries := database.New(db)

	u, err := queries.GetUserByScreenName(ctx, s)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGetUserByScreenNameUserNotFound
		}

		return nil, fmt.Errorf("queries.GetUserByScreenName: %w", err)
	}

	return &u, nil
}
