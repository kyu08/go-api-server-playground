package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	input := usecase.NewCreateUserInput(req.GetScreenName(), req.GetUserName(), req.GetBio())
	output, err := s.CreateUserUsecase.Run(ctx, input)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &api.CreateUserResponse{
		Id: output.ID.String(),
	}, nil
}
