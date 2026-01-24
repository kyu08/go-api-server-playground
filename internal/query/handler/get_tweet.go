package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/query/usecase"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetTweet(ctx context.Context, req *api.GetTweetRequest) (*api.GetTweetResponse, error) {
	input := usecase.NewTweetGetInput(req.GetTweetId())
	output, err := h.TweetGetUsecase.Run(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.GetTweetResponse{
		TweetId:           output.TweetID,
		Body:              output.Body,
		AuthorId:          output.AuthorID,
		AuthorScreenName:  output.AuthorScreenName,
		AuthorDisplayName: output.AuthorDisplayName,
		CreatedAt:         timestamppb.New(output.CreatedAt),
		UpdatedAt:         timestamppb.New(output.UpdatedAt),
	}, nil
}
