package handler

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/kyu08/go-api-server-playground/internal/grpcutil"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

const bufSize = 1024 * 1024

func setupTestServer(t *testing.T) (api.TwitterServiceClient, func()) {
	t.Helper()

	ctx := context.Background()

	client, dbTeardown, err := database.NewEmulatorWithClient(ctx)
	if err != nil {
		t.Fatalf("failed to create Spanner client: %v", err)
	}

	lis := bufconn.Listen(bufSize)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpcutil.ConversionError(),
		grpcutil.LoggerForTest(t),
	))

	twitterServer := NewTwitterServer(client)

	api.RegisterTwitterServiceServer(server, twitterServer)

	go func() {
		if err := server.Serve(lis); err != nil {
			t.Logf("server exited: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		ctx,
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
		dbTeardown()
	}

	return api.NewTwitterServiceClient(conn), cleanup
}
