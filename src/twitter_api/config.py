"""アプリケーション設定 (pydantic-settings)."""

from __future__ import annotations

from functools import lru_cache
from typing import Literal

from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    """環境変数から読み込まれるアプリケーション設定."""

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        env_prefix="APP_",
        extra="ignore",
    )

    # 例: sqlite+aiosqlite:///./app.db / postgresql+asyncpg://user:pass@host/db
    database_url: str = "sqlite+aiosqlite:///./app.db"
    # echo するかどうか。本番では False。
    database_echo: bool = False

    log_level: Literal["DEBUG", "INFO", "WARNING", "ERROR"] = "INFO"
    log_json: bool = True

    host: str = "0.0.0.0"  # noqa: S104 (ローカル検証用なので bind any で問題なし)
    port: int = 8080


@lru_cache(maxsize=1)
def get_settings() -> Settings:
    """設定をキャッシュして返す."""
    return Settings()
