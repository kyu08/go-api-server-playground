package handler

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/kyu08/go-api-server-playground/internal/grpcutil"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

const bufSize = 1024 * 1024

func setupTestServer(t *testing.T) (api.TwitterServiceClient, func()) {
	t.Helper()

	lis := bufconn.Listen(bufSize)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpcutil.ConversionError(),
		grpcutil.LoggerForTest(t),
	))

	twitterServer, teardown, err := NewTwitterServer()
	if err != nil {
		t.Fatalf("failed to create TwitterServer: %v", err)
	}

	api.RegisterTwitterServiceServer(server, twitterServer)

	go func() {
		if err := server.Serve(lis); err != nil {
			t.Logf("server exited: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return lis.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}

	cleanup := func() {
		conn.Close()
		server.Stop()
		teardown()
	}

	return api.NewTwitterServiceClient(conn), cleanup
}
