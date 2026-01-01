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
| Goコードのlint                 | [golangci-lint](https://github.com/golangci/golangci-lint), [gocapsule](https://github.com/YuitoSato/gocapsule) |

# 各種手順

## エンドポイント更新時の手順
1. `proto/twitter.proto`を更新
1. `make gen-proto`

## ローカルでの開発手順
1. `make dev-tools`で必要なツールをインストール
1. 必要に応じて`make test`, `make lint-go`, `make lint-gocapsule`などを実行

## gocapsuleについて
[gocapsule](https://github.com/YuitoSato/gocapsule)はカプセル化を強制するGo linterです。
`New**`コンストラクタが存在する構造体に対して、直接的な構造体リテラルの作成やフィールドの再代入を防ぎます。

### 実行方法
```bash
# ローカルで実行
make lint-gocapsule

# CIで自動実行されます
```
