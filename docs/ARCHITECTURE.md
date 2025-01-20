## ディレクトリ構成
```
./
├── cmd
│   └── server
├── docs                    // ドキュメント
├── e2e                     // E2Eテストのシナリオファイル、setup/teardownスクリプト
├── internal
│   ├── config              // 環境変数など
│   ├── domain              // ドメイン層
│   │   ├── entity          // ドメインモデル
│   │   ├── repository      // repository（のinterface）
│   │   └── service         // ドメインサービス
│   ├── errors              // アプリケーション内で共通のエラー構造体とメソッドの定義
│   ├── handler             // handler層
│   ├── infrastructure      // インフラ層
│   │   └── database
│   │       ├── docker_init // docker-composeでMySQLを起動する際の初期化スクリプト
│   │       ├── query       // 発行したいクエリを定義(sqlcを利用しているためクエリを独立したsqlファイルに記述する必要がある)
│   │       ├── repository  // repositoryの実装
│   │       └── schema      // MySQLのスキーマ定義
│   └── usecase             // usecase層
└── pkg
    └── api                 // protobuf定義
```

## 各層の責務と特徴
### handler
- アプリケーションの外側から渡ってきたデータをusecaseの都合に変換する
- usecaseから返ってきたresult, errorをアプリケーションの外側の都合に変換する

### usecase
- domain層に定義したビジネスロジックを呼び出してユースケースを達成する

### domain
- ドメインモデルの振る舞いを定義
- ビジネスロジックを定義
- repositoryのI/Fを定義
- 他の層に依存しない

### infrastructure
- 永続化層（今回だとMySQL）とのやりとり方法を規定

## 各層の依存関係
- `handler` → `usecase` → `domain`
- `infra` → `domain`

のような依存方向になっている
