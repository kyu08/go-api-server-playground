package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kyu08/go-api-server-playground/internal/config"
	pb "github.com/kyu08/go-api-server-playground/pkg/grpc"
	oursql "github.com/kyu08/go-api-server-playground/sql"
)

func (s *TwitterServer) Health(ctx context.Context, _ *pb.HealthRequest) (*pb.HealthResponse, error) {
	a := testSQL(ctx, s.config)

	return &pb.HealthResponse{Message: "twitter" + fmt.Sprintf("%+v", a)}, nil
}

func testSQL(ctx context.Context, config *config.Config) []oursql.Author {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	//nolint:exhaustruct,exhaustivestruct // 必要なフィールドだけ初期化したい
	mysqlConf := mysql.Config{
		User:             config.DBUser,
		Passwd:           config.DBPasswd,
		Addr:             config.DBAddr,
		DBName:           config.DBName,
		Net:              "tcp",
		Collation:        "utf8mb4_general_ci",
		Loc:              jst,
		MaxAllowedPacket: 0,
		ServerPubKey:     "",
		TLSConfig:        "",
		Logger:           nil,
		ParseTime:        true,
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		panic(err)
	}

	queries := oursql.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("authors: %v\n", authors)

	return authors
}
