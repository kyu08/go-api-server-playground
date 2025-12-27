package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/apstndb/spanemuboost"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kyu08/go-api-server-playground/internal/grpcutil"
	"github.com/kyu08/go-api-server-playground/internal/handler"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/pkg/api"
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

	twitterServer := handler.NewTwitterServer(client)

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
