package handler

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kyu08/go-api-server-playground/proto/api"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	client, cleanup := setupTestServer(t)
	defer cleanup()

	ctx := context.Background()

	// テストデータ
	screenName := "test_user"
	userName := "Test User"
	bio := "This is a bio"

	t.Run("正常にユーザーを作成できる", func(t *testing.T) {
		resp, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: screenName,
			UserName:   userName,
			Bio:        bio,
		})

		require.NoError(t, err)
		require.Len(t, resp.GetId(), 36) // UUID形式

		// 作成したユーザーを取得して確認
		findResp, err := client.FindUserByScreenName(ctx, &api.FindUserByScreenNameRequest{
			ScreenName: screenName,
		})

		require.NoError(t, err)
		require.Equal(t, screenName, findResp.GetScreenName())
		require.Equal(t, userName, findResp.GetUserName())
		require.Equal(t, bio, findResp.GetBio())
	})

	t.Run("screen_nameが空の場合エラー", func(t *testing.T) {
		_, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: "",
			UserName:   "Test User",
			Bio:        "bio",
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "screen name is required")
	})

	t.Run("user_nameが空の場合エラー", func(t *testing.T) {
		_, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: "s",
			UserName:   "",
			Bio:        "bio",
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "user name is required")
	})

	t.Run("userのコンストラクタでエラーが変える場合は場合エラー", func(t *testing.T) {
		_, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: strings.Repeat("a", 21),
			UserName:   "Test User",
			Bio:        "bio",
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "screen_name is too long")
	})

	t.Run("すでに存在するscreen_nameの場合エラー", func(t *testing.T) {
		// 同じscreen_nameで再度作成を試みる
		_, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: screenName,
			UserName:   "Second User",
			Bio:        "bio",
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "the screen name specified is already used")
	})
}
