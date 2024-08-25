package handler

import (
	"context"
	"fmt"

	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	input := usecase.NewCreateUserInput(req.GetScreenName(), req.GetUserName(), req.GetBio())
	output, err := s.CreateUserUsecase.Run(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("uc.Run: %w", err)
	}

	return &api.CreateUserResponse{
		Id: output.ID.String(),
	}, nil
}
