package user

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestNewUserBio(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in         string
		want       Bio
		wantErrMsg *string
	}{
		"validな引数を指定するとエラーが返らず、Bioが返る": {
			in:         "bio",
			want:       Bio("bio"),
			wantErrMsg: nil,
		},
		"invalidな引数を指定するとエラーと空文字が返る": {
			in:         "",
			want:       "",
			wantErrMsg: lo.ToPtr("bio is required"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUserBio(tt.in)

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

func TestBio_validate(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in         Bio
		wantErrMsg *string
	}{
		"1~160文字だとエラーが返らない(1文字)": {
			in:         "s",
			wantErrMsg: nil,
		},
		"1~160文字だとエラーが返らない(160文字)": {
			in:         Bio(strings.Repeat("s", 160)),
			wantErrMsg: nil,
		},
		"空文字だとエラーが返る": {
			in:         "",
			wantErrMsg: lo.ToPtr("bio is required"),
		},
		"161文字以上だとエラーが返る": {
			in:         Bio(strings.Repeat("s", 161)),
			wantErrMsg: lo.ToPtr("bio is too long"),
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
