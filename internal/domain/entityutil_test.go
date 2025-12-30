package domain

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

type TestUser struct{}

func TestNewID(t *testing.T) {
	t.Parallel()

	firstGeneration := NewID[TestUser]()
	t.Run("生成される文字数が36文字であること", func(t *testing.T) {
		t.Parallel()
		require.Len(t, firstGeneration.String(), 36)
	})

	t.Run("もう一度生成しても一致しないこと", func(t *testing.T) {
		t.Parallel()
		require.NotEqual(t, firstGeneration.String(), NewID[TestUser]().String())
	})
}

func TestNewFromString(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in           string
		expect       *ID[TestUser]
		expectErrMsg *string
	}{
		"uuid文字列を与えると正常に初期化できる": {
			in:           "f1ec0404-189b-49fa-a77c-37f8ff3c39b8",
			expect:       &ID[TestUser]{value: lo.Must(uuid.Parse("f1ec0404-189b-49fa-a77c-37f8ff3c39b8"))},
			expectErrMsg: nil,
		},
		"空文字を渡すとエラーが返る": {
			in:           "",
			expect:       nil,
			expectErrMsg: lo.ToPtr("invalid UUID length: 0"),
		},
		"uuidv4ではない文字列を渡すとエラーが返る": {
			in:           "zzz",
			expect:       nil,
			expectErrMsg: lo.ToPtr("invalid UUID length: 3"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := NewFromString[TestUser](tt.in)
			opt := []cmp.Option{cmpopts.IgnoreFields(ID[TestUser]{}, "_"), cmp.AllowUnexported(ID[TestUser]{})}

			if tt.expectErrMsg == nil {
				if diff := cmp.Diff(*tt.expect, got, opt...); diff != "" {
					t.Errorf("mismatch. (-expect +got)\n%s", diff)
				}
			} else {
				require.Equal(t, *tt.expectErrMsg, err.Error())
			}
		})
	}
}
