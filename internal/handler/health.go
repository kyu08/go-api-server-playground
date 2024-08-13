package handler

import (
	"context"
	"log"

	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func (s *TwitterServer) Health(ctx context.Context, _ *api.HealthRequest) (*api.HealthResponse, error) {
	log.Printf("Received: %v", "Health")

	return &api.HealthResponse{Message: "twitter"}, nil
}
