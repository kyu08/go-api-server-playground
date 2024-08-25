package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) Health(ctx context.Context, _ *api.HealthRequest) (*api.HealthResponse, error) {
	return &api.HealthResponse{Message: "twitter"}, nil
}
