package handler

import (
	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

type TwitterServer struct {
	pb.UnimplementedTwitterServiceServer
}

func NewTwitterServer() *TwitterServer {
	//nolint:exhaustruct,exhaustivestruct // 明示的に初期化する必要が特にない
	return &TwitterServer{}
}
