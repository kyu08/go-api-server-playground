package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kyu08/go-api-server-playground/internal/config"
	"github.com/kyu08/go-api-server-playground/internal/handler"
	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.New(context.Background())
	if err != nil {
		panic(err)
	}

	const (
		// docker composeで起動する際にhostを指定してしまうとうまく接続できないので空文字にしている。
		// ローカルでも起動したい場合は環境変数等で分岐するといいかもしれない
		host = ""
		port = 8080
	)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterTwitterServiceServer(server, handler.NewTwitterServer(config))

	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server on port %d", port)

		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
