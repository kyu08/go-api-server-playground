package grpcutil

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
)

func TestConvertErrorToGRPCStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		inputErr     error
		expectedCode codes.Code
		expectedMsg  string
	}{
		{
			name:         "NotFoundエラーはNotFoundに変換される",
			inputErr:     apperrors.NewNotFoundError("user"),
			expectedCode: codes.NotFound,
			expectedMsg:  "user not found",
		},
		{
			name:         "Preconditionエラーはinvalid argumentに変換される",
			inputErr:     apperrors.NewPreconditionError("invalid input"),
			expectedCode: codes.InvalidArgument,
			expectedMsg:  "invalid input",
		},
		{
			name:         "InternalエラーはInternalに変換される",
			inputErr:     apperrors.NewInternalError(errors.New("db connection failed")),
			expectedCode: codes.Internal,
			expectedMsg:  "internal server error",
		},
		{
			name:         "未知のエラーはInternalに変換される",
			inputErr:     errors.New("unknown error"),
			expectedCode: codes.Internal,
			expectedMsg:  "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := convertErrorToGRPCStatus(tt.inputErr)

			st, ok := status.FromError(result)
			if !ok {
				t.Fatalf("expected gRPC status error, got %v", result)
			}

			if st.Code() != tt.expectedCode {
				t.Errorf("expected code %v, got %v", tt.expectedCode, st.Code())
			}

			if st.Message() != tt.expectedMsg {
				t.Errorf("expected message %q, got %q", tt.expectedMsg, st.Message())
			}
		})
	}
}
