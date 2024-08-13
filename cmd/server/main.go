package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kyu08/go-api-server-playground/database"
	"github.com/kyu08/go-api-server-playground/internal/config"
	"github.com/kyu08/go-api-server-playground/internal/handler"
	"github.com/kyu08/go-api-server-playground/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.New(context.Background())
	if err != nil {
		panic(err)
	}

	db, err := database.NewDBConnection(config)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	api.RegisterTwitterServiceServer(server, handler.NewTwitterServer(db))

	reflection.Register(server)

	// TODO: アプリケーションのpanicをcatchする
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
