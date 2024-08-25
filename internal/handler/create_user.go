package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	log.Printf("Received: %v", "CreateUser") // TODO: インターセプター側でログ出力するようにする

	input := usecase.NewCreateUserInput(req.GetScreenName(), req.GetUserName(), req.GetBio())
	output, err := s.CreateUserUsecase.Run(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("uc.Run: %w", err)
	}

	return &api.CreateUserResponse{
		Id: output.ID.String(),
	}, nil
}
