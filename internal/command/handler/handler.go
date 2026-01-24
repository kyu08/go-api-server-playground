package handler

import (
	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/command/domain/user"
	"github.com/kyu08/go-api-server-playground/internal/command/infrastructure"
	"github.com/kyu08/go-api-server-playground/internal/command/usecase"
)

type Handler struct {
	TweetCreateUsecase *usecase.TweetCreateUsecase
	CreateUserUsecase  *usecase.CreateUserUsecase
}

func NewHandler(client *spanner.Client) *Handler {
	// repository実装
	tweetRepository := infrastructure.NewTweetRepository()
	userRepository := infrastructure.NewUserRepository()

	return &Handler{
		TweetCreateUsecase: usecase.NewTweetCreateUsecase(client, tweetRepository, userRepository),
		CreateUserUsecase:  usecase.NewCreateUserUsecase(client, user.NewCreateUserService(userRepository)),
	}
}
