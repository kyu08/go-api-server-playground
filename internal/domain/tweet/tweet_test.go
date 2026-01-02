package tweet

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestNewTweet(t *testing.T) {
	t.Parallel()

	authorID := domain.NewID[user.User]()

	type args struct {
		authorID domain.ID[user.User]
		body     string
	}

	tests := map[string]struct {
		in         args
		want       *Tweet
		wantErrMsg *string
	}{
		"validなparamsを渡すとTweetが返る": {
			in: args{
				authorID: authorID,
				body:     "hello world",
			},
			want: &Tweet{
				AuthorID: authorID,
				body:     "hello world",
			},
			wantErrMsg: nil,
		},
		"invalidなbodyを渡すとエラーが返る(空文字)": {
			in: args{
				authorID: authorID,
				body:     "",
			},
			want:       nil,
			wantErrMsg: lo.ToPtr("body is required"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewTweet(tt.in.authorID, tt.in.body)

			if tt.wantErrMsg != nil {
				require.Error(t, err)
				require.Equal(t, *tt.wantErrMsg, err.Error())
			} else {
				require.NoError(t, err)

				opts := []cmp.Option{
					cmpopts.IgnoreFields(Tweet{}, "ID", "AuthorID", "CreatedAt", "UpdatedAt"),
					cmp.AllowUnexported(Tweet{}),
				}
				if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
					t.Errorf("NewTweet() mismatch (-want +got) = \n%s", diff)
				}

				require.Equal(t, tt.in.authorID.String(), got.AuthorID.String())
			}
		})
	}
}

func TestTweet_NewFromDTO(t *testing.T) {
	t.Parallel()

	type args struct {
		id        string
		authorID  string
		body      string
		createdAt time.Time
		updatedAt time.Time
	}

	tests := map[string]struct {
		in      args
		want    *Tweet
		wantErr error
	}{
		"validなparamsを渡すとTweetが返る": {
			in: args{
				id:        "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				authorID:  "a541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				body:      "hello world",
				createdAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			want: &Tweet{
				ID:        lo.Must(domain.NewFromString[Tweet]("f541dd5f-48fd-40d6-8e5b-6c25b4681f3c")),
				AuthorID:  lo.Must(domain.NewFromString[user.User]("a541dd5f-48fd-40d6-8e5b-6c25b4681f3c")),
				body:      "hello world",
				CreatedAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		"invalidなidを渡すとエラーが返る": {
			in: args{
				id:        "",
				authorID:  "a541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				body:      "hello world",
				createdAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("invalid UUID length: 0"),
		},
		"invalidなauthorIDを渡すとエラーが返る": {
			in: args{
				id:        "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				authorID:  "",
				body:      "hello world",
				createdAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: errors.New("invalid UUID length: 0"),
		},
		"invalidなbodyを渡すとエラーが返る": {
			in: args{
				id:        "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				authorID:  "a541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				body:      "",
				createdAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: ErrBodyRequired,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewFromDTO(tt.in.id, tt.in.authorID, tt.in.body, tt.in.createdAt, tt.in.updatedAt)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)

				opts := []cmp.Option{
					cmpopts.IgnoreFields(Tweet{}, "ID", "AuthorID", "CreatedAt", "UpdatedAt"),
					cmp.AllowUnexported(Tweet{}),
				}
				if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
					t.Errorf("NewFromDTO() mismatch (-want +got) = \n%s", diff)
				}

				// IDとAuthorIDの文字列比較
				require.Equal(t, tt.in.id, got.ID.String())
				require.Equal(t, tt.in.authorID, got.AuthorID.String())
			}
		})
	}
}
