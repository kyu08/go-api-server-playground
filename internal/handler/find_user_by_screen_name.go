package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/proto/api"
)

func (s *TwitterServer) FindUserByScreenName(
	ctx context.Context,
	req *api.FindUserByScreenNameRequest,
) (*api.FindUserByScreenNameResponse, error) {
	input := usecase.NewFindUserByScreenNameInput(req.GetScreenName())

	output, err := s.FindUserByScreenNameUsecase.Run(ctx, input)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &api.FindUserByScreenNameResponse{
		Id:         output.ID,
		ScreenName: output.ScreenName,
		UserName:   output.UserName,
		Bio:        output.Bio,
	}, nil
}
