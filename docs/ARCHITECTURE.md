# アーキテクチャ

Python (FastAPI) における一般的な Web API の組み方を踏襲した、シンプルなレイヤ構成にしている.
旧 Go 実装の「ドメインモデル + CQRS + レイヤードアーキテクチャ」のような重厚な構成は採らない.

## 依存方向

```
routers  →  services  →  models / db
    ↓
schemas (Pydantic)
```

- **routers**: FastAPI のエンドポイント定義. リクエスト/レスポンスのスキーマ変換と、`Depends` でのサービス注入に専念する.
- **schemas (Pydantic)**: HTTP の境界での DTO. バリデーションは Pydantic のフィールド制約で行う.
- **services**: ビジネスロジック. ORM を直接触ることもあるが、router からは隠蔽する.
- **models (SQLAlchemy)**: テーブルの構造そのもの. `Base.metadata` から `alembic` がマイグレーションを生成する.
- **db**: エンジン / セッションファクトリの単一の置き場所. `deps.get_session` で `Depends` 経由で渡す.
- **exceptions**: 業務エラーは `AppError` のサブクラスとして送出し、FastAPI の例外ハンドラで HTTP ステータスにマップする.
- **config / logging**: 環境変数を `pydantic-settings` で型付き読み込み、`structlog` で JSON ログ.

## なぜ「ドメイン層 + CQRS」をやめたか

- Python の小〜中規模 Web API ではオーバーキルになりがちで、メタプログラミング・型システムとの相性も悪い.
- Pydantic + SQLAlchemy で「入力バリデーション」「永続化モデル」「業務ロジック」を分けるのが Python 慣習に近い.
- このリポジトリの目的は「Python で Web API を書く感覚を掴むこと」なので、慣習に寄せる方を優先した.

## エラーハンドリング

- 業務エラー (`NotFoundError` / `ConflictError` / `ValidationError`) は service 層から送出.
- FastAPI の例外ハンドラ (`exceptions.register_exception_handlers`) が `{"error": "...", "message": "..."}` の JSON に変換し、適切な HTTP ステータスを返す.
- リクエストスキーマ違反は Pydantic / FastAPI 標準の `422 Unprocessable Entity` がそのまま返る.

## DB マイグレーション

- `alembic` を使い、`models/` を真として `make migration MSG="..."` で差分を autogenerate.
- 開発の起動時 (`main._lifespan`) では SQLite に対し `Base.metadata.create_all` を呼んでテーブルを生成しているので、最初の起動だけは alembic を回さなくても動く.

## テスト方針

- ハンドラ層から DB までを通した API テストを `httpx.ASGITransport` + in-memory SQLite で行う.
- 各テストごとに独立した in-memory DB を `conftest.py` で立てる.
- 単体テストはサービス層のロジックが複雑になった際にスポット的に追加する想定.
