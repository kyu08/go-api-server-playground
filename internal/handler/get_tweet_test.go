package handler

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/kyu08/go-api-server-playground/proto/api"
)

func TestGetTweet(t *testing.T) {
	t.Parallel()

	client, cleanup := setupTestServer(t)
	defer cleanup()

	ctx := context.Background()

	t.Run("存在するTweetの詳細を取得できる", func(t *testing.T) {
		screenName := randomScreenName(t)
		createUserResp, err := client.CreateUser(ctx, &api.CreateUserRequest{
			ScreenName: screenName,
			UserName:   "昼寝のプロ",
			Bio:        "bio",
		})
		require.NoError(t, err)

		// tweetを作成
		createTweetResp, err := client.CreateTweet(ctx, &api.CreateTweetRequest{
			AuthorId: createUserResp.GetId(),
			Body:     "おひるねなう",
		})
		require.NoError(t, err)

		// 取得できるか確認
		tweetDetail, err := client.GetTweet(ctx, &api.GetTweetRequest{
			TweetId: createTweetResp.GetId(),
		})
		require.NoError(t, err)
		opt := []cmp.Option{
			protocmp.Transform(),
			protocmp.IgnoreFields(&api.GetTweetResponse{}, "created_at", "updated_at"),
		}
		expect := &api.GetTweetResponse{
			TweetId:           createTweetResp.GetId(),
			Body:              "おひるねなう",
			AuthorId:          createUserResp.GetId(),
			AuthorScreenName:  screenName,
			AuthorDisplayName: "昼寝のプロ",
		}
		if diff := cmp.Diff(expect, tweetDetail, opt...); diff != "" {
			t.Errorf("mismatch. (-expect +got)\n%s", diff)
		}
	})

	t.Run("TweetIDに空文字を指定するとエラーが返る", func(t *testing.T) {
		tweetDetail, err := client.GetTweet(ctx, &api.GetTweetRequest{
			TweetId: "",
		})
		require.Error(t, err)
		require.Nil(t, tweetDetail)
		assertGRPCError(t, err, codes.InvalidArgument, "tweet_id is required")
	})

	t.Run("不正な形式のTweetIDを指定するとエラーが返る", func(t *testing.T) {
		tweetDetail, err := client.GetTweet(ctx, &api.GetTweetRequest{
			TweetId: "invalid uuid",
		})
		require.Error(t, err)
		require.Nil(t, tweetDetail)
		assertGRPCError(t, err, codes.InvalidArgument, "invalid UUID length: 12")
	})

	t.Run("存在しないTweetIDを指定するとエラーが返る", func(t *testing.T) {
		tweetDetail, err := client.GetTweet(ctx, &api.GetTweetRequest{
			TweetId: "f6628467-5f01-47b7-8dba-896a723efb00",
		})
		require.Error(t, err)
		require.Nil(t, tweetDetail)
		assertGRPCError(t, err, codes.NotFound, "TweetDetail not found")
	})
}
