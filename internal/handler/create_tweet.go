package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/proto/api"
)

func (s *TwitterServer) CreateTweet(ctx context.Context, req *api.CreateTweetRequest) (*api.CreateTweetResponse, error) {
	input := usecase.NewTweetCreateInput(req.GetAuthorId(), req.GetBody())

	output, err := s.TweetCreateUsecase.Run(ctx, input)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &api.CreateTweetResponse{
		Id: output.ID.String(),
	}, nil
}
