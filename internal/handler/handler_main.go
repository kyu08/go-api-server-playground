package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

type TwitterServer struct {
	api.UnimplementedTwitterServiceServer

	CreateUserUsecase           *usecase.CreateUserUsecase
	FindUserByScreenNameUsecase *usecase.FindUserByScreenNameUsecase
}

func NewTwitterServer() (*TwitterServer, func(), error) {
	ctx := context.Background()

	// NOTE: このプロジェクトはあくまでアプリケーションアーキテクチャ検証用のプロジェクトなのでローカルでしか起動しない。
	// そのためエミュレーターに接続する前提で実装している。
	client, teardown, err := database.NewEmulatorWithClient(ctx)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	userRepository := repository.NewUserRepository()

	return &TwitterServer{
		UnimplementedTwitterServiceServer: api.UnimplementedTwitterServiceServer{},
		CreateUserUsecase:                 usecase.NewCreateUserUsecase(client, userRepository),
		FindUserByScreenNameUsecase:       usecase.NewFindUserByScreenNameUsecase(client, userRepository),
	}, teardown, nil
}
