package server

import (
	"context"

	"cloud.google.com/go/spanner"
	commandhandler "github.com/kyu08/go-api-server-playground/internal/command/handler"
	queryhandler "github.com/kyu08/go-api-server-playground/internal/query/handler"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
)

// TwitterServer は command と query のハンドラーを統合するサーバー
type TwitterServer struct {
	api.UnimplementedTwitterServiceServer

	CommandHandler *commandhandler.Handler
	QueryHandler   *queryhandler.Handler
}

func NewTwitterServer(client *spanner.Client) *TwitterServer {
	return &TwitterServer{
		UnimplementedTwitterServiceServer: api.UnimplementedTwitterServiceServer{},
		CommandHandler:                    commandhandler.NewHandler(client),
		QueryHandler:                      queryhandler.NewHandler(client),
	}
}

// Command methods
func (s *TwitterServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	return s.CommandHandler.CreateUser(ctx, req)
}

func (s *TwitterServer) CreateTweet(ctx context.Context, req *api.CreateTweetRequest) (*api.CreateTweetResponse, error) {
	return s.CommandHandler.CreateTweet(ctx, req)
}

// Query methods
func (s *TwitterServer) FindUserByScreenName(
	ctx context.Context,
	req *api.FindUserByScreenNameRequest,
) (*api.FindUserByScreenNameResponse, error) {
	return s.QueryHandler.FindUserByScreenName(ctx, req)
}

func (s *TwitterServer) GetTweet(ctx context.Context, req *api.GetTweetRequest) (*api.GetTweetResponse, error) {
	return s.QueryHandler.GetTweet(ctx, req)
}

func (s *TwitterServer) Health(ctx context.Context, req *api.HealthRequest) (*api.HealthResponse, error) {
	return s.QueryHandler.Health(ctx, req)
}
