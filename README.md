# go-api-server-playground
[kyu08](https://github.com/kyu08)がGoの素振りをするためのリポジトリです。

# 実装内容や使用技術など

| 題材                           | Twitter風のAPIサーバー                                     |
| :---                           | :---                                                       |
| 言語                           | Go 1.22                                                    |
| 通信方式                       | gRPC                                                       |
| DB                             | MySQL 8.4                                                  |
| CI                             | GitHub Actions                                             |
| 依存関係更新                   | dependabot                                                 |
| ローカルでのコンテナ実行ツール | docker compose                                             |
| E2Eテストツール                | [runn](https://github.com/k1LoW/runn)                      |
| Goコードのlint                 | [golangci-lint](https://github.com/golangci/golangci-lint) |
| SQLからのコード生成            | [sqlc](https://github.com/sqlc-dev/sqlc)                   |
| SQLのlint, format              | [sqlfluff](https://github.com/sqlfluff/sqlfluff)           |

# 各種手順

## エンドポイント更新時の手順
1. `api/twitter.proto`を更新
1. `make gen-proto`

## DDL更新時の手順
1. `./sql/schema/schema.sql`を更新
1. `make gen-sqlc`

## DML更新時の手順
1. `./sql/query/query.sql`を更新
1. `make gen-sqlc`

## ローカルでの起動手順
1. `make container-up`でコンテナを起動する

## ローカルでの開発手順
1. `make dev-tools`で必要なツールをインストール
1. 必要に応じて`make test-e2e`, `make lint`などを実行
