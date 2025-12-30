package domain

import (
	"testing"

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
