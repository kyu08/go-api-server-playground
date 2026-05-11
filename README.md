# go-api-server-playground

> 旧 Go 実装を Python (FastAPI) にリプレイスしたリポジトリ。

[kyu08](https://github.com/kyu08) が Python で API Server を書くときの構成・ライブラリ選定・周辺ツールを検証するためのリポジトリです。
(リポジトリ名は歴史的経緯で `go-` を残していますが中身は Python です)

## 構成

| 題材           | Twitter 風の API サーバー                                                                |
| :---           | :---                                                                                       |
| 言語           | Python 3.12                                                                                |
| 通信方式       | HTTP / JSON (REST)                                                                         |
| Webフレームワーク | [FastAPI](https://fastapi.tiangolo.com/)                                                |
| ORM            | [SQLAlchemy 2.0 (async)](https://docs.sqlalchemy.org/en/20/)                              |
| バリデーション | [Pydantic v2](https://docs.pydantic.dev/)                                                  |
| 設定           | [pydantic-settings](https://docs.pydantic.dev/latest/concepts/pydantic_settings/)          |
| マイグレーション | [Alembic](https://alembic.sqlalchemy.org/)                                              |
| DB             | SQLite (dev/test) / PostgreSQL (prod 想定)                                                 |
| 依存関係管理   | [uv](https://docs.astral.sh/uv/) + `pyproject.toml`                                        |
| Lint / Format  | [Ruff](https://docs.astral.sh/ruff/)                                                       |
| 型チェック     | [mypy --strict](https://mypy.readthedocs.io/)                                              |
| テスト         | [pytest](https://docs.pytest.org/) + `pytest-asyncio` + `httpx.ASGITransport`              |
| ログ           | [structlog](https://www.structlog.org/) (JSON 出力)                                        |
| CI             | GitHub Actions                                                                             |
| 依存関係更新   | dependabot                                                                                 |

## ディレクトリ構成

```text
.
├── pyproject.toml          # 依存・ツール設定の単一窓口
├── uv.lock                 # 依存解決のロック
├── Makefile                # よく使うタスク
├── alembic.ini             # Alembic 設定
├── migrations/             # DB マイグレーション
│   ├── env.py
│   └── versions/
├── src/
│   └── twitter_api/        # アプリ本体 (src layout)
│       ├── __init__.py
│       ├── __main__.py     # `python -m twitter_api` のエントリ
│       ├── main.py         # FastAPI app factory
│       ├── config.py       # pydantic-settings
│       ├── deps.py         # FastAPI Depends ヘルパ
│       ├── exceptions.py   # アプリ例外と handler
│       ├── logging.py      # structlog 設定
│       ├── db/             # DB エンジン / Base
│       ├── models/         # SQLAlchemy ORM モデル
│       ├── schemas/        # Pydantic スキーマ (DTO)
│       ├── services/       # ビジネスロジック
│       └── routers/        # FastAPI Router
└── tests/                  # pytest テスト
    ├── conftest.py
    ├── test_health.py
    ├── test_users.py
    └── test_tweets.py
```

Python の Web API としては、routers → services → models (SQLAlchemy) というシンプルなフローを採っています。
Pydantic スキーマがリクエスト/レスポンスの境界を担い、`Depends` でセッションとサービスを注入します。

## エンドポイント

| Method | Path                     | 概要                                |
| :----- | :----------------------- | :---------------------------------- |
| GET    | `/healthz`               | ヘルスチェック                      |
| POST   | `/users`                 | ユーザー作成                        |
| GET    | `/users/{screen_name}`   | screen_name でユーザー取得          |
| POST   | `/tweets`                | ツイート作成                        |
| GET    | `/tweets/{tweet_id}`     | ツイート詳細取得                    |

OpenAPI ドキュメントは起動後 `http://localhost:8080/docs` で確認できます。

## セットアップ

[uv](https://docs.astral.sh/uv/) がインストール済みであることを前提とします。

```sh
# 依存を入れる (.venv が自動で作られる)
make install

# 開発サーバ (自動リロード)
make dev

# OpenAPI を見る
open http://localhost:8080/docs
```

環境変数は `APP_` プレフィックスで設定できます。`.env` も読み込みます。

| 変数                 | デフォルト                              | 説明                                |
| :------------------- | :-------------------------------------- | :---------------------------------- |
| `APP_DATABASE_URL`   | `sqlite+aiosqlite:///./app.db`          | SQLAlchemy URL                      |
| `APP_DATABASE_ECHO`  | `false`                                 | SQL を stdout に出すか              |
| `APP_LOG_LEVEL`      | `INFO`                                  | ログレベル                          |
| `APP_LOG_JSON`       | `true`                                  | JSON ログ vs 人間向けカラーログ     |
| `APP_HOST`           | `0.0.0.0`                               | バインドする host                   |
| `APP_PORT`           | `8080`                                  | バインドする port                   |

## よく使うコマンド

```sh
make install      # uv sync
make dev          # 開発サーバ (auto reload)
make run          # 通常起動
make test         # pytest
make test-cov     # pytest + coverage
make lint         # ruff check
make format       # ruff format + ruff check --fix
make typecheck    # mypy --strict
make check        # lint + typecheck + test
make migrate      # alembic upgrade head
make migration MSG="add column foo"   # 新規 migration を autogenerate
```

## 依存管理の方針

- **`uv` 一本化**: 仮想環境 (`.venv`)、Python バージョン (`.python-version`)、依存解決 (`uv.lock`) すべてを `uv` で管理。
- **`pyproject.toml` 単一ファイル**: 依存・Ruff・mypy・pytest・coverage の設定をすべてここに集約。
- **`[project.dependencies]`** が本番、**`[dependency-groups].dev`** が開発のみの依存。
- **`uv.lock` は必ずコミット**。CI でも `uv sync --frozen` で再現性を担保。
- **依存更新は dependabot**。`uv` ecosystem (lockfile 更新) と `github-actions` を monthly でまとめ更新。
