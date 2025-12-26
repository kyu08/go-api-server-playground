package handler

import (
	"context"
	"net"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/kyu08/go-api-server-playground/internal/errors"
	"github.com/kyu08/go-api-server-playground/pkg/api"
)

const bufSize = 1024 * 1024

func setupTestServer(t *testing.T) (api.TwitterServiceClient, func()) {
	t.Helper()

	lis := bufconn.Listen(bufSize)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		conversionErrorInterceptor(),
		loggerInterceptor(t),
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

func conversionErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			if errors.IsNotFound(err) {
				return resp, status.Error(codes.NotFound, err.Error())
			}

			if errors.IsPrecondition(err) {
				return resp, status.Error(codes.InvalidArgument, err.Error())
			}

			return resp, status.Error(codes.Internal, "internal server error")
		}

		return resp, err
	}
}

func loggerInterceptor(t *testing.T) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		methodName := strings.Split(info.FullMethod, "/")[2]

		t.Logf("[gRPC] start: %s, request: %+v", methodName, req)

		resp, err := handler(ctx, req)
		if err != nil {
			t.Logf("[gRPC] error: %s, error: %v, stack: %s", methodName, err, errors.GetStackTrace(err))
		} else {
			t.Logf("[gRPC] end: %s, response: %+v", methodName, resp)
		}

		return resp, err
	}
}
