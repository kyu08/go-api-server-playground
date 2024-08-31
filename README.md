# 実装内容や使用技術など

| 言語                           | Go 1.23                                                    |
| :---                           | :---                                                       |
| 通信方式                       | REST                                                       |
| DB                             | MySQL 8.4                                                  |
| CI                             | GitHub Actions                                             |
| 依存関係更新                   | dependabot                                                 |
| ローカルでのコンテナ実行ツール | docker compose                                             |
| E2Eテストツール                | [runn](https://github.com/k1LoW/runn)                      |
| Goコードのlint                 | [golangci-lint](https://github.com/golangci/golangci-lint) |
| SQLからのコード生成            | [sqlc](https://github.com/sqlc-dev/sqlc)                   |
| SQLのlint, format              | [sqlfluff](https://github.com/sqlfluff/sqlfluff)           |

# 各種手順

## DDL更新時の手順
1. `./infrastructure/database/schema/schema.sql`を更新
1. `make gen-sqlc`

## DML更新時の手順
1. `./infrastructure/database/query/query.sql`を更新
1. `make gen-sqlc`

## E2Eテストの実行手順
1. `make test-e2e-with-refresh`でコンテナを起動しE2Eテストを実行

## ローカルでの開発手順
1. `make dev-tools`で必要なツールをインストール
1. 必要に応じて`make test-e2e`, `make lint`などを実行
