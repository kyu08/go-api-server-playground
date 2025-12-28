package handler

import (
	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

type TwitterServer struct {
	api.UnimplementedTwitterServiceServer

	CreateUserUsecase           *usecase.CreateUserUsecase
	FindUserByScreenNameUsecase *usecase.FindUserByScreenNameUsecase
}

func NewTwitterServer(client *spanner.Client) *TwitterServer {
	userRepository := repository.NewUserRepository()

	return &TwitterServer{
		UnimplementedTwitterServiceServer: api.UnimplementedTwitterServiceServer{},
		CreateUserUsecase:                 usecase.NewCreateUserUsecase(client, user.NewUserService(userRepository)),
		FindUserByScreenNameUsecase:       usecase.NewFindUserByScreenNameUsecase(client, userRepository),
	}
}
