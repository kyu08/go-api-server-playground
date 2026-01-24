package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"cloud.google.com/go/spanner"
	"github.com/apstndb/spanemuboost"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	commandhandler "github.com/kyu08/go-api-server-playground/internal/command/handler"
	queryhandler "github.com/kyu08/go-api-server-playground/internal/query/handler"
	"github.com/kyu08/go-api-server-playground/internal/shared/grpcutil"
	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// NOTE: このプロジェクトはあくまでアプリケーションアーキテクチャ検証用のプロジェクトなのでローカルでしか起動しない。
	// そのためエミュレーターに接続する前提で実装している。
	emulator, emulatorTeardown, err := spanemuboost.NewEmulator(context.Background(), spanemuboost.EnableInstanceAutoConfigOnly())
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer emulatorTeardown()

	client, teardown, err := database.GetSpannerClient(emulator)
	if err != nil {
		panic(err)
	}
	defer teardown()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpcutil.ConversionError(),
		grpcutil.Logger(logger),
		grpc_recovery.UnaryServerInterceptor(),
	))

	twitterServer := NewTwitterServer(client)

	api.RegisterTwitterServiceServer(server, twitterServer)
	reflection.Register(server)

	go func() {
		const (
			// NOTE: docker composeで起動する際にhostを指定してしまうとうまく接続できないので空文字にしている。
			// ローカルでも起動したい場合は環境変数等で分岐するといいかもしれない(起動はできるが毎回プロンプトが表示されて面倒)
			host = ""
			port = 8080
		)

		logger.Info(fmt.Sprintf("start gRPC server on port %d", port))

		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			panic(err)
		}

		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping gRPC server...")
	server.GracefulStop() // NOTE: 受け付けているリクエストを捌き切ってからサーバーを停止するために必要
}

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
func (s *TwitterServer) FindUserByScreenName(ctx context.Context, req *api.FindUserByScreenNameRequest) (*api.FindUserByScreenNameResponse, error) {
	return s.QueryHandler.FindUserByScreenName(ctx, req)
}

func (s *TwitterServer) GetTweet(ctx context.Context, req *api.GetTweetRequest) (*api.GetTweetResponse, error) {
	return s.QueryHandler.GetTweet(ctx, req)
}

func (s *TwitterServer) Health(ctx context.Context, req *api.HealthRequest) (*api.HealthResponse, error) {
	return s.QueryHandler.Health(ctx, req)
}
