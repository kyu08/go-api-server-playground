package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kyu08/go-api-server-playground/internal/config"
	"github.com/kyu08/go-api-server-playground/internal/errors"
)

//nolint:mnd //ここではマジックナンバーの用途は見ればわかるので許容
func NewDBConnection(config *config.Config) (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, errors.WithStack(errors.NewInternalError(err))
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
		return nil, errors.WithStack(errors.NewInternalError(err))
	}

	// NOTE: 下記は仮置きの数字。パフォーマンス要件に応じて調整する必要がある。
	db.SetMaxOpenConns(100)                 // DBへの最大接続数
	db.SetMaxIdleConns(50)                  // アイドル状態の最大コネクション数
	db.SetConnMaxLifetime(10 * time.Minute) // コネクションの最大ライフタイム
	db.SetConnMaxIdleTime(5 * time.Minute)  // コネクションの最大アイドル時間

	return db, nil
}
