package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) FindUserByScreenName(
	ctx context.Context,
	req *api.FindUserByScreenNameRequest,
) (*api.FindUserByScreenNameResponse, error) {
	log.Printf("Received: %v", "FindUserByScreenName") // TODO: インターセプター側でログ出力するようにする

	screenName, err := user.NewUserScreenName(req.GetScreenName())
	if err != nil {
		return nil, fmt.Errorf("user.NewUserScreenName: %w", err)
	}

	u, err := s.userRepository.FindByScreenName(ctx, s.db, string(screenName))
	if err != nil {
		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return &api.FindUserByScreenNameResponse{
		Id:         u.ID.String(),
		ScreenName: u.ScreenName.String(),
		UserName:   u.UserName.String(),
		Bio:        u.Bio.String(),
	}, nil
}
