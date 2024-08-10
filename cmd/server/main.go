package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	host := "127.0.0.1"
	port := 8080

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterTwitterServiceServer(s, NewTwitterServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server on port %d", port)
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

type twitterServer struct {
	pb.UnimplementedTwitterServiceServer
}

func NewTwitterServer() *twitterServer {
	return &twitterServer{}
}

// TODO: 実装例があれば参考にしつつ別パッケージに分離する
func (s *twitterServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Message: "twitter"}, nil
}
