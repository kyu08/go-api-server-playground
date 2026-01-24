package handler

import (
	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/query/infrastructure"
	"github.com/kyu08/go-api-server-playground/internal/query/usecase"
)

type Handler struct {
	TweetGetUsecase             *usecase.TweetGetUsecase
	FindUserByScreenNameUsecase *usecase.FindUserByScreenNameUsecase
}

func NewHandler(client *spanner.Client) *Handler {
	// query実装
	tweetQuery := infrastructure.NewTweetQuery()
	userQuery := infrastructure.NewUserQuery()

	return &Handler{
		TweetGetUsecase:             usecase.NewTweetGetUsecase(client, tweetQuery),
		FindUserByScreenNameUsecase: usecase.NewFindUserByScreenNameUsecase(client, userQuery),
	}
}
