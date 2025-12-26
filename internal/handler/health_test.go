package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kyu08/go-api-server-playground/pkg/api"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	client, cleanup := setupTestServer(t)
	defer cleanup()

	resp, err := client.Health(context.Background(), &api.HealthRequest{})

	require.NoError(t, err)
	require.Equal(t, "twitter", resp.Message)
}
