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

	// TODO: ログ、エラーハンドリング(アプリケーションのエラーから判断してステータスコードをいい感じにする)のインターセプタを追加する
	//nolint:lll //URLなので仕方なし
	// see: https://zenn.dev/jinn/articles/d3b177eafbc457#%E5%8D%98%E4%B8%80%E3%81%AE%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%82%BB%E3%83%97%E3%82%BF%E3%83%BC%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B%E5%A0%B4%E5%90%88
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
