package handler

import (
	"github.com/kyu08/go-api-server-playground/internal/config"
	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

type TwitterServer struct {
	pb.UnimplementedTwitterServiceServer
	config *config.Config
}

func NewTwitterServer(config *config.Config) *TwitterServer {
	//nolint:exhaustruct,exhaustivestruct // 明示的に初期化する必要が特にない
	return &TwitterServer{
		config: config,
	}
}
