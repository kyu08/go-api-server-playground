# go-api-server-playground
[kyu08](https://github.com/kyu08)がGoの素振りをするためのリポジトリです。

# 実装内容や使用技術など

| 題材                           | Twitter風のAPIサーバー                                     |
| :---                           | :---                                                       |
| 言語                           | Go 1.25.4                                                  |
| 通信方式                       | gRPC                                                       |
| DB                             | Cloud Spanner (Emulator)                                   |
| CI                             | GitHub Actions                                             |
| 依存関係更新                   | dependabot                                                 |
| Goコードのlint                 | [golangci-lint](https://github.com/golangci/golangci-lint) |

# 各種手順

## エンドポイント更新時の手順
1. `api/twitter.proto`を更新
1. `make gen-proto`

## ローカルでの開発手順
1. `make dev-tools`で必要なツールをインストール
1. 必要に応じて`make test`, `make lint`などを実行
