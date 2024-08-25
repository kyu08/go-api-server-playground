package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) FindUserByScreenName(
	ctx context.Context,
	req *api.FindUserByScreenNameRequest,
) (*api.FindUserByScreenNameResponse, error) {
	log.Printf("Received: %v", "FindUserByScreenName") // TODO: インターセプター側でログ出力するようにする

	input := usecase.NewFindUserByScreenNameInput(req.GetScreenName())
	output, err := s.FindUserByScreenNameUsecase.Run(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("uc.Run: %w", err)
	}

	return &api.FindUserByScreenNameResponse{
		Id:         output.ID.String(),
		ScreenName: output.ScreenName.String(),
		UserName:   output.UserName.String(),
		Bio:        output.Bio.String(),
	}, nil
}
