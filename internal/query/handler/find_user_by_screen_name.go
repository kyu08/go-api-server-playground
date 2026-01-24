package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/query/usecase"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
)

func (h *Handler) FindUserByScreenName(
	ctx context.Context,
	req *api.FindUserByScreenNameRequest,
) (*api.FindUserByScreenNameResponse, error) {
	input := usecase.NewFindUserByScreenNameInput(req.GetScreenName())

	output, err := h.FindUserByScreenNameUsecase.Run(ctx, input)
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
