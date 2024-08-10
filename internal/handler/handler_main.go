package handler

import (
	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

type TwitterServer struct {
	pb.UnimplementedTwitterServiceServer
}

func NewTwitterServer() *TwitterServer {
	return &TwitterServer{}
}
