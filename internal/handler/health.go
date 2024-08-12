package handler

import (
	"context"
	"log"

	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

func (s *TwitterServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	log.Printf("Received: %v", "Health")

	return &pb.HealthResponse{Message: "twitter"}, nil
}
