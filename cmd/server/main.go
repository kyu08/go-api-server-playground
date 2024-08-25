package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/internal/handler"
	"github.com/kyu08/go-api-server-playground/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// TODO: エラーハンドリング(アプリケーションのエラーから判断してステータスコードをいい感じにする)のインターセプタを追加する
	// TODO: アプリケーションのpanicをcatchしてinternal server errorを返すようなインターセプタを追加する

	server := grpc.NewServer(grpc.UnaryInterceptor(loggerInterceptor()))
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

		log.Printf("start gRPC server on port %d", port)

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
	log.Println("stopping gRPC server...")
	server.GracefulStop() // NOTE: 受け付けているリクエストを捌き切ってからサーバーを停止するために必要
}

func loggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("[start]: %s(%s)", strings.Split(info.FullMethod, "/")[2], req)
		defer log.Printf("[end]:   %s(%s)", strings.Split(info.FullMethod, "/")[2], req)

		resp, err := handler(ctx, req)
		if err != nil {
			if !errors.IsPreconditionError(err) {
				// TODO: ここでスタックトレースをログ出力する
				log.Printf("[error]: %s(internal: %s)", strings.Split(info.FullMethod, "/")[2], err)
				return resp, errors.NewInternalError()
			}
			log.Printf("[error]: %s(%s)", strings.Split(info.FullMethod, "/")[2], err)
		}

		return resp, err
	}
}
