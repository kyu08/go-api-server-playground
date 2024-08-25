package user

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
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
				ID:         "",
				ScreenName: "screen_name",
				UserName:   "user_name",
				Bio:        "bio",
				IsPrivate:  false,
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
			tt := tt
			t.Parallel()
			got, err := New(tt.in.screenName, tt.in.userName, tt.in.bio)

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
