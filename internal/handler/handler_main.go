package handler

import (
	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database/repository"
	"github.com/kyu08/go-api-server-playground/internal/usecase"
	"github.com/kyu08/go-api-server-playground/proto/api"
)

type TwitterServer struct {
	api.UnimplementedTwitterServiceServer

	TweetCreateUsecase          *usecase.TweetCreateUsecase
	TweetGetUsecase             *usecase.TweetGetUsecase
	CreateUserUsecase           *usecase.CreateUserUsecase
	FindUserByScreenNameUsecase *usecase.FindUserByScreenNameUsecase
}

func NewTwitterServer(client *spanner.Client) *TwitterServer {
	tweetRepository := repository.NewTweetRepository()
	userRepository := repository.NewUserRepository()

	return &TwitterServer{
		UnimplementedTwitterServiceServer: api.UnimplementedTwitterServiceServer{},
		TweetCreateUsecase:                usecase.NewTweetCreateUsecase(client, tweetRepository, userRepository),
		TweetGetUsecase:                   usecase.NewTweetGetUsecase(client, tweetRepository, userRepository),
		CreateUserUsecase:                 usecase.NewCreateUserUsecase(client, user.NewUserService(userRepository)),
		FindUserByScreenNameUsecase:       usecase.NewFindUserByScreenNameUsecase(client, userRepository),
	}
}
