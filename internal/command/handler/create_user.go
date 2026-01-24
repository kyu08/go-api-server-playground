package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/command/usecase"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
)

func (h *Handler) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	input := usecase.NewCreateUserInput(req.GetScreenName(), req.GetUserName(), req.GetBio())

	output, err := h.CreateUserUsecase.Run(ctx, input)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &api.CreateUserResponse{
		Id: output.ID,
	}, nil
}
