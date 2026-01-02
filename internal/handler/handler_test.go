package handler

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/apstndb/spanemuboost"
	"github.com/google/uuid"
	"github.com/kyu08/go-api-server-playground/internal/grpcutil"
	"github.com/kyu08/go-api-server-playground/internal/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/proto/api"
	"github.com/stretchr/testify/require"
	tcspanner "github.com/testcontainers/testcontainers-go/modules/gcloud/spanner"
)

const (
	bufSize    = 1024 * 1024
	uuidLength = 36
)

var spannerEmulator *tcspanner.Container

func TestMain(m *testing.M) {
	emulator, emulatorTeardown, err := spanemuboost.NewEmulator(context.Background(), spanemuboost.EnableInstanceAutoConfigOnly())
	if err != nil {
		log.Fatalln(err)
		return
	}

	spannerEmulator = emulator
	exitCode := m.Run()

	// TestMainはm.Run()の戻り値を使ってos.Exitを呼び出す必要がある。（そうしないとテスト失敗時にプロセスがexitCode: 0で終了してしまい、
	// テストが成功したとみなされてしまう。
	// defer emulatorTeardown()を使う前提だと別途関数を切らないとうまく書けないのでdeferを使わずにここで明示的に呼び出している。
	emulatorTeardown()
	os.Exit(exitCode)
}

// setupTestServer テスト用のDBとgGRPCサーバーを立ち上げる。
func setupTestServer(t *testing.T) (api.TwitterServiceClient, func()) {
	t.Helper()
	client, teardown, err := database.GetSpannerClient(spannerEmulator)
	if err != nil {
		t.Fatalf("failed to get spanner client: %s", err)
	}

	lis := bufconn.Listen(bufSize)
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpcutil.ConversionError(),
		grpcutil.LoggerForTest(t),
	))

	twitterServer := NewTwitterServer(client)
	api.RegisterTwitterServiceServer(server, twitterServer)

	go func() {
		if err := server.Serve(lis); err != nil {
			t.Logf("server exited: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return lis.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}

	cleanup := func() {
		_ = conn.Close()

		server.Stop()
		teardown()
	}

	return api.NewTwitterServiceClient(conn), cleanup
}

func assertGRPCError(t *testing.T, err error, wantCode codes.Code, wantMessage string) {
	t.Helper()

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, wantCode, st.Code())
	require.Contains(t, st.Message(), wantMessage)
}

// テスト用のscreen nameをランダムに生成して返す。
func randomScreenName(t *testing.T) string {
	t.Helper()
	return uuid.New().String()[:20]
}
