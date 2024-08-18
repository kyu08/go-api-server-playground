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

var ErrFindUserByScreenNameUserNotFound = errors.New("user not found")

func (s *TwitterServer) FindUserByScreenName(
	ctx context.Context,
	req *api.FindUserByScreenNameRequest,
) (*api.FindUserByScreenNameResponse, error) {
	log.Printf("Received: %v", "FindUserByScreenName") // TODO: インターセプター側でログ出力するようにする

	screenName, err := user.NewUserScreenName(req.GetScreenName())
	if err != nil {
		return nil, fmt.Errorf("user.NewUserScreenName: %w", err)
	}

	u, err := FindUserByScreenName(ctx, s.db, string(screenName))
	if err != nil {
		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return &api.FindUserByScreenNameResponse{
		Id:         u.ID,
		ScreenName: u.ScreenName,
		UserName:   u.UserName,
		Bio:        u.Bio,
	}, nil
}

// TODO: パッケージ構成をいい感じにする
// 1. repositoryパターン
func FindUserByScreenName(ctx context.Context, db *sql.DB, s string) (*database.User, error) {
	queries := database.New(db)

	u, err := queries.FindUserByScreenName(ctx, s)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrFindUserByScreenNameUserNotFound
		}

		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return &u, nil
}
