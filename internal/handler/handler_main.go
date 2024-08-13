package handler

import (
	"database/sql"

	"github.com/kyu08/go-api-server-playground/pkg/api"
)

type TwitterServer struct {
	api.UnimplementedTwitterServiceServer
	db *sql.DB
}

func NewTwitterServer(db *sql.DB) *TwitterServer {
	//nolint:exhaustruct,exhaustivestruct // 明示的に初期化する必要が特にない
	return &TwitterServer{
		db: db,
	}
}
