package tweet

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestBody_validate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in         Body
		wantErrMsg *string
	}{
		"1~140文字だとエラーが返らない(1文字)": {
			in:         "s",
			wantErrMsg: nil,
		},
		"1~140文字だとエラーが返らない(140文字)": {
			in:         Body(strings.Repeat("s", 140)),
			wantErrMsg: nil,
		},
		"141文字以上だとエラーが返る": {
			in:         Body(strings.Repeat("s", 141)),
			wantErrMsg: lo.ToPtr("body is too long"),
		},
		"1~140文字だとエラーが返らない(140文字,マルチバイト文字)": {
			in:         Body(strings.Repeat("あ", 140)),
			wantErrMsg: nil,
		},
		"141文字以上だとエラーが返る(マルチバイト文字)": {
			in:         Body(strings.Repeat("あ", 141)),
			wantErrMsg: lo.ToPtr("body is too long"),
		},
		"空文字だとエラーが返る": {
			in:         "",
			wantErrMsg: lo.ToPtr("body is required"),
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

func TestNewBody(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in         string
		want       Body
		wantErrMsg *string
	}{
		"validな引数を指定するとエラーが返らず、Bodyが返る": {
			in:         "body",
			want:       lo.Must(NewBody("body")),
			wantErrMsg: nil,
		},
		"invalidな引数を指定するとエラーと空文字が返る": {
			in:         "",
			want:       "",
			wantErrMsg: lo.ToPtr("body is required"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewBody(tt.in)

			if tt.wantErrMsg != nil {
				require.Error(t, err)
				require.Equal(t, *tt.wantErrMsg, err.Error())
			} else {
				require.NoError(t, err)
			}

			if got != tt.want {
				t.Errorf("NewBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
