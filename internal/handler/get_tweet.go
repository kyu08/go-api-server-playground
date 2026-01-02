package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/proto/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *TwitterServer) GetTweet(ctx context.Context, req *api.GetTweetRequest) (*api.GetTweetResponse, error) {
	input := usecase.NewTweetGetInput(req.GetTweetId())

	output, err := s.TweetGetUsecase.Run(ctx, input)
	if err != nil {
		return nil, apperrors.WithStack(err)
	}

	return &api.GetTweetResponse{
		TweetId:           output.TweetID.String(),
		Body:              output.Body,
		AuthorId:          output.AuthorId.String(),
		AuthorScreenName:  output.AuthorScreenName,
		AuthorDisplayName: output.AuthorDisplayName,
		CreatedAt:         timestamppb.New(output.CreatedAt),
		UpdatedAt:         timestamppb.New(output.UpdatedAt),
	}, nil
}
