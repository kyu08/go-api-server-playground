package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/command/usecase"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
)

func (h *Handler) CreateTweet(ctx context.Context, req *api.CreateTweetRequest) (*api.CreateTweetResponse, error) {
	input := usecase.NewTweetCreateInput(req.GetAuthorId(), req.GetBody())

	output, err := h.TweetCreateUsecase.Run(ctx, input)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &api.CreateTweetResponse{
		Id: output.ID,
	}, nil
}
