package handler

import (
	"database/sql"

	"github.com/kyu08/go-api-server-playground/internal/database/repository"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

type TwitterServer struct {
	api.UnimplementedTwitterServiceServer
	db             *sql.DB
	userRepository *repository.UserRepository // NOTE: テスト時にrepositoryをモック化せずに本物のDBを使うつもりなのでinterfaceではなく実体に依存している。
}

func NewTwitterServer(db *sql.DB, userRepository *repository.UserRepository) *TwitterServer {
	//nolint:exhaustruct,exhaustivestruct // 明示的に初期化する必要が特にない
	return &TwitterServer{
		db:             db,
		userRepository: userRepository,
	}
}
