.PHONY: install run dev test test-cov lint format typecheck check ci migrate migration clean

# =========================================
# 開発環境構築
# =========================================
install: ## 依存をインストール (uv が必要)
	uv sync

# =========================================
# アプリケーションの起動
# =========================================
run: ## サーバを起動 (uvicorn)
	uv run twitter-api

dev: ## 開発モード起動 (自動リロード)
	uv run uvicorn twitter_api.main:app --reload --host 0.0.0.0 --port 8080

# =========================================
# テスト
# =========================================
test: ## pytest 実行
	uv run pytest

test-cov: ## カバレッジ付きで pytest 実行
	uv run pytest --cov --cov-report=term-missing

# =========================================
# 静的解析
# =========================================
lint: ## ruff lint
	uv run ruff check .

format: ## ruff format + lint --fix
	uv run ruff format .
	uv run ruff check . --fix

typecheck: ## mypy 実行
	uv run mypy src tests

check: lint typecheck test ## lint + typecheck + test

ci: check ## CI 用 (check と同じ)

# =========================================
# DB マイグレーション (alembic)
# =========================================
migrate: ## 最新へマイグレーション
	uv run alembic upgrade head

migration: ## 新規 migration 生成 (例: make migration MSG="add foo")
	uv run alembic revision --autogenerate -m "$(MSG)"

# =========================================
# Misc
# =========================================
clean: ## キャッシュ削除
	rm -rf .pytest_cache .mypy_cache .ruff_cache .coverage htmlcov
	find . -type d -name __pycache__ -exec rm -rf {} + 2>/dev/null || true

help: ## このヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?##' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*##"}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
