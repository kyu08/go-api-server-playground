package handler

import (
	"database/sql"

	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

type TwitterServer struct {
	pb.UnimplementedTwitterServiceServer
	db *sql.DB
}

func NewTwitterServer(db *sql.DB) *TwitterServer {
	//nolint:exhaustruct,exhaustivestruct // 明示的に初期化する必要が特にない
	return &TwitterServer{
		db: db,
	}
}
