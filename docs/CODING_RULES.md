# コーディング規約
このドキュメントでは、本リポジトリで採用している特徴的なコーディング規約を記載する。

## Linter設定
golangci-lintで`default: all`を採用し、必要に応じて個別のlinterを無効化している。

これはgolangci-lintのバージョンアップ時に新規追加されたlinterを試したうえで有効化するかどうかを判断したいためである。

## エラーハンドリング
- 自リポジトリ内のパッケージで発生したエラーは発生場所でラップする。（スタックトレースを付与するため）
    ```go
    import (
        "github.com/kyu08/go-api-server-playground/internal/apperrors"
    )
    func someFunc() error {
        // ...

        isExisting, err := s.IsExistingScreenName(ctx, rtx, user.ScreenName)
        if err != nil {
            return apperrors.WithStack(err)
        }

        // ...
    }
    ```

## コードフォーマット
- gofumpt, goimportsを使用(golangci-lintを使ってCIで実行しているためCIが通れば問題ない)

## handler, usecase層のファイル名の命名規則
- handler層のファイル名はRPC名をスネークケースに変換したものとする。
    - 例: CreateUser RPCのhandler層のファイル名は`create_user.go`
- usecase層のファイル名は`${エンティティ名}_${操作}.go`の形式とする。
    - 例: CreateUserユースケースのusecase層のファイル名は`user_create.go`
- 背景
    - handle層のファイルはRPCと一対一に対応しているためRPC名をファイル名にすることで対応関係が明確になる。
    - 一方でusecase層のファイルは上記の形式にしておくことで、ファイル一覧を見た際にエンティティごとのどのようなユースケースが存在するかを把握しやすくなる利点があると考えこのような規約を設定している。
