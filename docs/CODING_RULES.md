# コーディング規約

## Lint / Format
- `ruff` を `select = ["ALL"]` で運用し、必要なルールだけ `ignore` する方針 (旧 Go 実装の golangci-lint と同じ思想).
- 新しい ruff にルールが追加されたら、まず試してから採用可否を判断する.
- フォーマッタは `ruff format` (Black 互換).

## 型
- `mypy --strict` を CI で実行.
- 関数は基本的に型注釈を付ける. 例外はテストの fixture 等.

## 命名
- ルーティングは REST 流の名詞ベース (`/users`, `/tweets`).
- テストは `tests/test_<対象>.py` に置き、関数名は `test_xxx` または日本語可 (per-file-ignore で許可).
- Pydantic スキーマは `XxxRequest` / `XxxResponse` を基本とする.

## バリデーション
- HTTP 入力のバリデーションは **Pydantic スキーマ** に書く. service 層の入り口では信頼してよい.
- 業務ロジック由来のバリデーション (重複, 存在チェック等) は **service 層で AppError サブクラスを送出**.

## エラー
- 業務エラーは `twitter_api.exceptions.AppError` のサブクラス (`NotFoundError` / `ConflictError` / `ValidationError`) を `raise`.
- これらは FastAPI の例外ハンドラで HTTP ステータスに変換される.
- 個別の HTTP ステータスを返したい場合は router 内で `HTTPException` を直接送出してもよい.

## 非同期
- service / router / DB アクセスはすべて `async def`.
- 同期コードと混ぜたい場合は `anyio.to_thread.run_sync` でラップする.

## 依存管理
- 依存追加は `uv add <package>` (本番) または `uv add --dev <package>` (開発のみ).
- `uv.lock` は必ずコミット. CI では `uv sync --frozen`.
