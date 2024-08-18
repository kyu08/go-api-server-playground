package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) CreateUser(
	ctx context.Context,
	req *api.CreateUserRequest,
) (*api.CreateUserResponse, error) {
	log.Printf("Received: %v", "CreateUser") // TODO: インターセプター側でログ出力するようにする

	user, err := user.New(req.ScreenName, req.UserName, req.Bio)
	if err != nil {
		return nil, fmt.Errorf("user.NewUser: %w", err)
	}

	if err := s.userRepository.Create(ctx, s.db, user); err != nil {
		return nil, fmt.Errorf("queries.FindUserByScreenName: %w", err)
	}

	return &api.CreateUserResponse{}, nil
}
