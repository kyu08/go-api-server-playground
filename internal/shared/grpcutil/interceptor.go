package grpcutil

import (
	"context"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
)

func ConversionError() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, convertErrorToGRPCStatus(err)
		}

		return resp, err
	}
}

//nolint:wrapcheck // gRPCステータスエラーはラップせずそのまま返す
func convertErrorToGRPCStatus(err error) error {
	if apperrors.IsNotFound(err) {
		return status.Error(codes.NotFound, err.Error())
	}

	if apperrors.IsPrecondition(err) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return status.Error(codes.Internal, "internal server error")
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
