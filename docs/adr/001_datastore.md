---
status: 'approved'
---

# 決定
データストアとして **SQLite (dev/test) + PostgreSQL (prod 想定)** を採用し、ORM は **SQLAlchemy 2.0 (async)** を使う.

# 文脈
- このリポジトリの目的は Python Web API の構成検証であり、データストア選定そのものは検証対象ではない.
- Python Web API のコミュニティ慣習では SQLAlchemy + RDBMS が圧倒的に主流.
- 旧 Go 実装は Cloud Spanner (Emulator) を使っていたが、Python では `google-cloud-spanner` の使い勝手と Web API 慣習の乖離が大きいため見送り.

# 論点
- SQLAlchemy か、より軽量な ORM (Tortoise, SQLModel, Piccolo 等) か
  - SQLAlchemy が de facto standard. async 対応も成熟しており、Alembic との組み合わせで migration まで賄える.
- DB を SQLite で完結させるか、PostgreSQL を Docker で立てるか
  - 開発/テストは SQLite で十分軽い. 本番想定だけ PostgreSQL 用に `asyncpg` を依存に入れておく.

# 決定時に想定した影響
- `make dev` だけでローカル動作する (DB の事前準備不要).
- SQLite と PostgreSQL の差異 (JSON 型, 並行性, etc.) は将来的に問題になりうる. 必要になったら CI に PostgreSQL ジョブを足す.

# 参考
- SQLAlchemy 2.0 async docs: https://docs.sqlalchemy.org/en/20/orm/extensions/asyncio.html
