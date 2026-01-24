package testutil

import (
	"context"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/spanner"
	"github.com/apstndb/spanemuboost"
	"github.com/kyu08/go-api-server-playground/internal/shared/apperrors"
	"github.com/kyu08/go-api-server-playground/internal/shared/grpcutil"
	"github.com/kyu08/go-api-server-playground/internal/shared/infrastructure/database"
	"github.com/kyu08/go-api-server-playground/internal/shared/proto/api"
	tcspanner "github.com/testcontainers/testcontainers-go/modules/gcloud/spanner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize = 1024 * 1024

	// UUIDLength はUUIDの文字列長
	UUIDLength = 36
)

var spannerEmulator *tcspanner.Container

// SetupTestMain はTestMainで呼び出し、エミュレーターを起動してteardown関数を返す
func SetupTestMain(m *testing.M) {
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

// ServerFactory はSpannerクライアントを受け取り、TwitterServiceServerを返す関数型
type ServerFactory func(client *spanner.Client) api.TwitterServiceServer

// SetupTestServer テスト用のDBとgRPCサーバーを立ち上げる。
// serverFactory はSpannerクライアントを受け取り、TwitterServiceServerを返す関数
func SetupTestServer(t *testing.T, serverFactory ServerFactory) (api.TwitterServiceClient, func()) {
	t.Helper()
	client, teardown, err := database.GetSpannerClient(spannerEmulator)
	if err != nil {
		t.Fatalf("failed to get spanner client: %s", err)
	}

	lis := bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		grpcutil.ConversionError(),
		loggerForTest(t),
	))

	twitterServer := serverFactory(client)
	api.RegisterTwitterServiceServer(grpcServer, twitterServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
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

		grpcServer.Stop()
		teardown()
	}

	return api.NewTwitterServiceClient(conn), cleanup
}

func loggerForTest(t *testing.T) grpc.UnaryServerInterceptor {
	t.Helper()
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		methodName := strings.Split(info.FullMethod, "/")[2]

		t.Logf("[gRPC] start: %s, request: %+v", methodName, req)

		resp, err := handler(ctx, req)
		if err != nil {
			t.Logf("[gRPC] error: %s, error: %v, stack: %s", methodName, err, apperrors.GetStackTrace(err))
		} else {
			t.Logf("[gRPC] end: %s, response: %+v", methodName, resp)
		}

		return resp, err
	}
}
