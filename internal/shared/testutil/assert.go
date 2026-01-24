package testutil

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AssertGRPCError はgRPCエラーのコードとメッセージを検証する
func AssertGRPCError(t *testing.T, err error, wantCode codes.Code, wantMessage string) {
	t.Helper()

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, wantCode, st.Code())
	require.Equal(t, wantMessage, st.Message())
}

// RandomScreenName はテスト用のscreen nameをランダムに生成して返す
func RandomScreenName(t *testing.T) string {
	t.Helper()
	return uuid.New().String()[:20]
}
