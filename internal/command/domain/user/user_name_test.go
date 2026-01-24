package user

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestUserName_validate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in         UserName
		wantErrMsg *string
	}{
		"1~20文字だとエラーが返らない(1文字)": {
			in:         "s",
			wantErrMsg: nil,
		},
		"1~20文字だとエラーが返らない(20文字)": {
			in:         UserName(strings.Repeat("s", 20)),
			wantErrMsg: nil,
		},
		"空文字だとエラーが返る": {
			in:         "",
			wantErrMsg: lo.ToPtr("user_name is required"),
		},
		"21文字以上だとエラーが返る": {
			in:         UserName(strings.Repeat("s", 21)),
			wantErrMsg: lo.ToPtr("user_name is too long"),
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

func TestNewUserUserName(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in         string
		want       UserName
		wantErrMsg *string
	}{
		"validな引数を指定するとエラーが返らず、UserNameが返る": {
			in:         "user_name",
			want:       UserName("user_name"),
			wantErrMsg: nil,
		},
		"invalidな引数を指定するとエラーと空文字が返る": {
			in:         "",
			want:       "",
			wantErrMsg: lo.ToPtr("user_name is required"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewUserUserName(tt.in)

			if tt.wantErrMsg != nil {
				require.Error(t, err)
				require.Equal(t, *tt.wantErrMsg, err.Error())
			} else {
				require.NoError(t, err)
			}

			if got != tt.want {
				t.Errorf("NewUserUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
