package user

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestScreenName_validate(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in         ScreenName
		wantErrMsg *string
	}{
		"1~20文字だとエラーが返らない(1文字)": {
			in:         "s",
			wantErrMsg: nil,
		},
		"1~20文字だとエラーが返らない(20文字)": {
			in:         ScreenName(strings.Repeat("s", 20)),
			wantErrMsg: nil,
		},
		"空文字だとエラーが返る": {
			in:         "",
			wantErrMsg: lo.ToPtr("screen_name is required"),
		},
		"21文字以上だとエラーが返る": {
			in:         ScreenName(strings.Repeat("s", 21)),
			wantErrMsg: lo.ToPtr("screen_name is too long"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := tt.in.validate()
			if tt.wantErrMsg != nil {
				require.Error(t, err)
				require.Equal(t, *tt.wantErrMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewUserScreenName(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in         string
		want       ScreenName
		wantErrMsg *string
	}{
		"validな引数を指定するとエラーが返らず、ScreenNameが返る": {
			in:         "screen_name",
			want:       ScreenName("screen_name"),
			wantErrMsg: nil,
		},
		"invalidな引数を指定するとエラーと空文字が返る": {
			in:         "",
			want:       "",
			wantErrMsg: lo.ToPtr("screen_name is required"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUserScreenName(tt.in)

			if tt.wantErrMsg != nil {
				require.Error(t, err)
				require.Equal(t, *tt.wantErrMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
