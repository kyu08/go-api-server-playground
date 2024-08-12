# go-api-server-playground
[kyu08](https://github.com/kyu08)がGoの素振りをするためのリポジトリです。

# 実装内容や使用技術など

| 題材                           | Twitter風のAPIサーバー                                     |
| ---                            | ---                                                        |
| 言語                           | Go                                                         |
| 通信方式                       | gRPC                                                       |
| DB                             | MySQL 8.4                                                  |
| CI                             | GitHub Actions                                             |
| E2Eテストツール                | [runn](https://github.com/k1LoW/runn)                      |
| 依存関係更新                   | dependabot                                                 |
| Linter                         | [golangci-lint](https://github.com/golangci/golangci-lint) |
| SQLからのコード生成            | [sqlc](https://github.com/sqlc-dev/sqlc)                   |
| ローカルでのコンテナ実行ツール | docker compose                                             |
| ローカルでの環境変数管理       | [direnv](https://github.com/direnv/direnv)                 |

# DDL更新時の手順
1. `./sql/schema/schema.sql`を更新
1. `make gen-sqlc`

# DML更新時の手順
1. `./sql/query/query.sql`を更新
1. `make gen-sqlc`
