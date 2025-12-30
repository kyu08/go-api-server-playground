package user

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	type args struct {
		screenName string
		userName   string
		bio        string
	}

	tests := map[string]struct {
		in         args
		want       *User
		wantErrMsg *string
	}{
		"validなUserを渡すとnilが返る": {
			in: args{
				screenName: "screen_name",
				userName:   "user_name",
				bio:        "bio",
			},
			want: &User{
				ID:         domain.NewID[User](),
				ScreenName: "screen_name",
				UserName:   "user_name",
				Bio:        "bio",
				CreatedAt:  time.Time{},
			},
			wantErrMsg: nil,
		},
		"invalidなUserを渡すとエラーが返る(ScreenName)": {
			in: args{
				screenName: "",
				userName:   "user_name",
				bio:        "bio",
			},
			want:       nil,
			wantErrMsg: lo.ToPtr("screen_name is required"),
		},
		"invalidなUserを渡すとエラーが返る(UserName)": {
			in: args{
				screenName: "screen_name",
				userName:   "",
				bio:        "bio",
			},
			want:       nil,
			wantErrMsg: lo.ToPtr("user_name is required"),
		},
		"invalidなUserを渡すとエラーが返る(Bio)": {
			in: args{
				screenName: "screen_name",
				userName:   "user_name",
				bio:        "",
			},
			want:       nil,
			wantErrMsg: lo.ToPtr("bio is required"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewUser(tt.in.screenName, tt.in.userName, tt.in.bio)

			if tt.wantErrMsg != nil {
				require.Error(t, err)
				require.Equal(t, *tt.wantErrMsg, err.Error())
			} else {
				require.NoError(t, err)
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(User{}, "ID", "CreatedAt")); diff != "" {
				t.Errorf("New() mismatch (-want +got) = \n%s", diff)
			}
		})
	}
}

func TestNewFromDTO(t *testing.T) {
	t.Parallel()

	type args struct {
		id         string
		screenName string
		userName   string
		bio        string
		createdAt  time.Time
	}

	tests := map[string]struct {
		in      args
		want    *User
		wantErr error
	}{
		"validなUserを渡すとnilが返る": {
			in: args{
				id:         "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				screenName: "screen_name",
				userName:   "user_name",
				bio:        "bio",
				createdAt:  time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			want: &User{
				ID:         lo.Must(domain.NewFromString[User]("f541dd5f-48fd-40d6-8e5b-6c25b4681f3c")),
				ScreenName: "screen_name",
				UserName:   "user_name",
				Bio:        "bio",
				CreatedAt:  time.Date(2025, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		"invalidなidを渡すとエラーが返る": {
			in: args{
				id:         "",
				screenName: "",
				userName:   "user_name",
				bio:        "bio",
			},
			want:    nil,
			wantErr: errors.New("invalid UUID length: 0"),
		},
		"invalidなscreenNameを渡すとエラーが返る": {
			in: args{
				id:         "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				screenName: "",
				userName:   "user_name",
				bio:        "bio",
			},
			want:    nil,
			wantErr: ErrScreenNameRequired,
		},
		"invalidなuserNameを渡すとエラーが返る": {
			in: args{
				id:         "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				screenName: "sn",
				userName:   "",
				bio:        "bio",
			},
			want:    nil,
			wantErr: ErrUserNameRequired,
		},
		"invalidなbioを渡すとエラーが返る": {
			in: args{
				id:         "f541dd5f-48fd-40d6-8e5b-6c25b4681f3c",
				screenName: "sn",
				userName:   "un",
				bio:        "",
			},
			want:    nil,
			wantErr: ErrBioRequired,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewFromDTO(tt.in.id, tt.in.screenName, tt.in.userName, tt.in.bio, tt.in.createdAt)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(User{}, "ID", "CreatedAt")); diff != "" {
					t.Errorf("NewFromDTO() mismatch (-want +got) = \n%s", diff)
				}
			}
		})
	}
}
