package handler

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	"github.com/kyu08/go-api-server-playground/proto/api"
)

func TestCreateTweet(t *testing.T) {
	t.Parallel()

	client, cleanup := setupTestServer(t)
	defer cleanup()

	ctx := context.Background()

	t.Run("正常にTweetを作成できる_140文字_マルチバイト文字", func(t *testing.T) {
		userResp, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: randomScreenName(t),
			UserName:   "Test User",
			Bio:        "bio",
		})
		require.NoError(t, err)

		// tweetを作成
		body := strings.Repeat("あ", 140)
		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: userResp.GetId(),
			Body:     body,
		})

		// 取得して確認
		require.NoError(t, err)
		require.Len(t, tweetResp.GetId(), uuidLength)

		tweetDetail, err := client.GetTweet(ctx, &api.GetTweetRequest{
			TweetId: tweetResp.GetId(),
		})
		require.NoError(t, err)
		require.Equal(t, tweetResp.GetId(), tweetDetail.GetTweetId())
		require.Equal(t, body, tweetDetail.GetBody())
	})

	t.Run("正常にTweetを作成できる_140文字_ASCII文字", func(t *testing.T) {
		userResp, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: randomScreenName(t),
			UserName:   "Test User",
			Bio:        "bio",
		})
		require.NoError(t, err)

		// tweetを作成
		body := strings.Repeat("あ", 140)
		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: userResp.GetId(),
			Body:     body,
		})

		require.NoError(t, err)
		require.Len(t, tweetResp.GetId(), uuidLength)

		// 取得して確認
		tweetDetail, err := client.GetTweet(ctx, &api.GetTweetRequest{
			TweetId: tweetResp.GetId(),
		})
		require.NoError(t, err)
		require.Equal(t, tweetResp.GetId(), tweetDetail.GetTweetId())
		require.Equal(t, body, tweetDetail.GetBody())
	})

	t.Run("AuthorIDに空文字を指定するとエラーが返る", func(t *testing.T) {
		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: "",
			Body:     "またーりついーとなう",
		})

		require.Nil(t, tweetResp)
		assertGRPCError(t, err, codes.InvalidArgument, "author_id is required")
	})

	t.Run("Bodyに空文字を指定するとエラーが返る", func(t *testing.T) {
		userResp, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: randomScreenName(t),
			UserName:   "Test User",
			Bio:        "bio",
		})
		require.NotNil(t, userResp)
		require.NoError(t, err)

		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: userResp.GetId(),
			Body:     "",
		})

		require.Nil(t, tweetResp)
		assertGRPCError(t, err, codes.InvalidArgument, "body is required")
	})

	t.Run("AuthorIDにuuid形式でない文字列を指定するとエラーが返る", func(t *testing.T) {
		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: "not-uuid",
			Body:     "またーりついーとなう",
		})

		require.Nil(t, tweetResp)
		assertGRPCError(t, err, codes.InvalidArgument, "invalid UUID length: 8")
	})

	t.Run("存在しないユーザーを指定するとエラーが返る", func(t *testing.T) {
		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: "f7fb794a-0b96-4807-8bff-14774d1adbce",
			Body:     "またーりついーとなう",
		})

		require.Nil(t, tweetResp)
		assertGRPCError(t, err, codes.NotFound, "user not found")
	})

	t.Run("NewTweetからエラーが返る引数を指定するとエラーが返る", func(t *testing.T) {
		userResp, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: randomScreenName(t),
			UserName:   "Test User",
			Bio:        "bio",
		})
		require.NotNil(t, userResp)
		require.NoError(t, err)

		tweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: userResp.GetId(),
			Body:     strings.Repeat("t", 141),
		})

		require.Nil(t, tweetResp)
		assertGRPCError(t, err, codes.InvalidArgument, "body is too long")
	})
}
