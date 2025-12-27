package grpcutil

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
)

func ConversionError() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			if apperrors.IsNotFound(err) {
				return resp, status.Error(codes.NotFound, err.Error())
			}

			if apperrors.IsPrecondition(err) {
				return resp, status.Error(codes.InvalidArgument, err.Error())
			}

			return resp, status.Error(codes.Internal, "internal server error")
		}

		return resp, err
	}
}

func Logger(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		methodName := strings.Split(info.FullMethod, "/")[2]

		logger.Info("start", slog.String("method", methodName), slog.Any("request", req))
		defer logger.Info("end", slog.String("method", methodName), slog.Any("request", req))

		resp, err := handler(ctx, req)
		if err != nil {
			if !apperrors.IsPrecondition(err) {
				logger.Error(err.Error(), "method", methodName, "error", apperrors.GetStackTrace(err))
			} else {
				logger.Warn(err.Error(), "method", methodName, "error", apperrors.GetStackTrace(err))
			}
		}

		return resp, err
	}
}

// TestLogger is a logger interface for testing.
type TestLogger interface {
	Logf(format string, args ...any)
}

func LoggerForTest(t TestLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		methodName := strings.Split(info.FullMethod, "/")[2]

		t.Logf("[gRPC] start: %s, request: %+v", methodName, req)

		resp, err := handler(ctx, req)
		if err != nil {
			t.Logf("[gRPC] error: %s, error: %v, stack: %s", methodName, err, apperrors.GetStackTrace(err))
		} else {
			t.Logf("[gRPC] end: %s, response: %+v", methodName, resp)
		}

		return resp, err
	}
}
