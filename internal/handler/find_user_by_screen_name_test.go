package handler

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func TestFindUserByScreenName(t *testing.T) {
	t.Parallel()

	client, cleanup := setupTestServer(t)
	defer cleanup()

	ctx := context.Background()

	// テストデータ
	screenName := "test_user"
	userName := "Test User"
	bio := "This is a bio"

	// まずユーザーを作成
	createResp, err := client.CreateUser(ctx, &api.CreateUserRequest{
		ScreenName: screenName,
		UserName:   userName,
		Bio:        bio,
	})
	require.NoError(t, err)

	t.Run("存在するユーザーを検索できる", func(t *testing.T) {
		resp, err := client.FindUserByScreenName(ctx, &api.FindUserByScreenNameRequest{
			ScreenName: screenName,
		})

		require.NoError(t, err)
		require.Equal(t, createResp.Id, resp.Id)
		require.Equal(t, screenName, resp.ScreenName)
		require.Equal(t, bio, resp.Bio)
	})

	t.Run("存在しないユーザーはエラー", func(t *testing.T) {
		_, err := client.FindUserByScreenName(ctx, &api.FindUserByScreenNameRequest{
			ScreenName: "nonexistent_user",
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.NotFound, st.Code())
	})
	t.Run("ScreenNameに空文字を指定するとエラー", func(t *testing.T) {
		_, err := client.FindUserByScreenName(ctx, &api.FindUserByScreenNameRequest{
			ScreenName: "",
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
	})
	t.Run("ScreenNameに21文字以上の文字列を指定するとエラー", func(t *testing.T) {
		_, err := client.FindUserByScreenName(ctx, &api.FindUserByScreenNameRequest{
			ScreenName: strings.Repeat("a", 21),
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
	})
}
