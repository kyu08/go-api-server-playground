package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *MarketServer) Health(ctx context.Context, _ *api.HealthRequest) (*api.HealthResponse, error) {
	return &api.HealthResponse{Message: "market"}, nil
}
