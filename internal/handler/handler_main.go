package handler

import (
	"context"

	"github.com/kyu08/go-api-server-playground/internal/config"
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

func NewTwitterServer() (*TwitterServer, error) {
	config, err := config.New(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db, err := database.NewDBConnection(config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userRepository := repository.NewUserRepository()

	return &TwitterServer{
		UnimplementedTwitterServiceServer: api.UnimplementedTwitterServiceServer{},
		CreateUserUsecase:                 usecase.NewCreateUserUsecase(db, userRepository),
		FindUserByScreenNameUsecase:       usecase.NewFindUserByScreenNameUsecase(db, userRepository),
	}, nil
}
