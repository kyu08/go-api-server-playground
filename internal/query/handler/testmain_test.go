package handler_test

import (
	"testing"

	"cloud.google.com/go/spanner"
	"github.com/kyu08/go-api-server-playground/internal/server"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
	"github.com/kyu08/go-api-server-playground/internal/shared/testutil"
)

func TestMain(m *testing.M) {
	testutil.SetupTestMain(m)
}

func setupTestServer(t *testing.T) (api.TwitterServiceClient, func()) {
	t.Helper()
	return testutil.SetupTestServer(t, func(client *spanner.Client) api.TwitterServiceServer {
		return server.NewTwitterServer(client)
	})
}
