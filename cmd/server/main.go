package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strings"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/handler"
	"github.com/kyu08/go-api-server-playground/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		loggerInterceptor(logger),
		grpc_recovery.UnaryServerInterceptor(),
	))
	twitterServer, err := handler.NewTwitterServer()
	if err != nil {
		panic(err)
	}

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

func loggerInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		methodName := strings.Split(info.FullMethod, "/")[2]

		logger.Info("start", slog.String("method", methodName), slog.Any("request", req))
		defer logger.Info("end", slog.String("method", methodName), slog.Any("request", req))

		resp, err := handler(ctx, req)
		if err != nil {
			logger.Error("error", "method", methodName, "error", errors.GetStackTrace(err))
			if !errors.IsPrecondition(err) {
				return resp, status.Error(codes.Internal, "internal server error")
			}
			return resp, status.Error(codes.InvalidArgument, err.Error())
		}

		return resp, err
	}
}
