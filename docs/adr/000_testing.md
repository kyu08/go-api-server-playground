---
status: 'approved'
---

# 概要
本ドキュメントはこのプロジェクトのテスト方針について記載した ADR である.

# 決定
- API テスト (router → service → DB) を `httpx.ASGITransport` + in-memory SQLite で行うことを基本とする.
- テスト DB は **テストごとに使い捨て** で立てる (`conftest.py` の `app` fixture).
- service 層のユニットテストは、ロジックが複雑な箇所にスポット的に追加する.
- Pydantic / SQLAlchemy のような外部 OSS の挙動はテストしない.

# 文脈
旧 Go 実装ではアーキテクチャ検証のため、E2E テスト中心の方針を採っていた. Python 実装でも同じ思想を引き継ぎつつ、Spanner Emulator のような重い依存は持たない方が Python の慣習に近い.

# 論点
- 単体テストか API テストか
  - API テスト中心. 内部実装をリファクタしても外形が変わらなければテストを直さなくていい.
- DB をモックするか実 DB を使うか
  - in-memory SQLite を使う. Docker / testcontainers が不要で、`uv run pytest` だけで完結する.

# 決定時に想定した影響
- リファクタリングしやすい (内部構造を変えても API テストは通る).
- テスト実行時間は十分速い (3 秒未満).
- SQLite と本番 (PostgreSQL) の方言差で取りこぼす可能性はあるが、最初は許容する.

# 参考
N/A
