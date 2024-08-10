package handler

import (
	"context"

	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

func (s *TwitterServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Message: "twitter"}, nil
}
