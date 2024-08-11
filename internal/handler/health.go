package handler

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
)

func (s *TwitterServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Message: "twitter"}, nil
}

func testSQL() {
	ctx := context.Background()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	const (
		User   = "user"
		Passwd = "password"
		Addr   = "localhost:3306"
		DBName = "db"
	)

	mysqlConf := mysql.Config{
		User:             User,
		Passwd:           Passwd,
		Addr:             Addr,
		DBName:           DBName,
		Net:              "tcp",
		Collation:        "utf8mb4_unicode_ci", // TODO: 調べる
		Loc:              jst,
		MaxAllowedPacket: 0,
		ServerPubKey:     "",
		TLSConfig:        "",
		Logger:           nil,
		ParseTime:        true,
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
}
