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

        isExisting, err := s.IsExistingScreenName(ctx, tx, user.ScreenName)
        if err != nil {
            return apperrors.WithStack(err)
        }

        // ...
    }
    ```

## コードフォーマット
- gofumpt, goimportsを使用(golangci-lintを使ってCIで実行しているためCIが通れば問題ない)
