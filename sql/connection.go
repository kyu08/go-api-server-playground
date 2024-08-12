package sql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kyu08/go-api-server-playground/internal/config"
)

func NewDBConnection(config *config.Config) (*sql.DB, error) {
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
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return db, nil
}