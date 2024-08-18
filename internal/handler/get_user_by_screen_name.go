package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

var ErrGetUserByScreenNameUserNotFound = errors.New("user not found")

func (s *TwitterServer) GetUserByScreenName(
	ctx context.Context,
	req *api.GetUserByScreenNameRequest,
) (*api.GetUserByScreenNameResponse, error) {
	log.Printf("Received: %v", "GetUserByScreenName")

	screenName, err := user.NewUserScreenName(req.GetScreenName())
	if err != nil {
		return nil, fmt.Errorf("user.NewUserScreenName: %w", err)
	}

	u, err := testSQL(ctx, s.db, string(screenName))
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
