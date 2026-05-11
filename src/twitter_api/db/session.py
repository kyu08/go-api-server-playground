"""SQLAlchemy エンジンとセッション管理.

FastAPI Web API の慣習に従い、エンジン/セッションファクトリはアプリの lifespan で 1 度だけ
作る. リクエストごとのセッションは ``deps.get_session`` で `Depends` 経由で渡す.
"""

from __future__ import annotations

from typing import TYPE_CHECKING

from sqlalchemy.ext.asyncio import (
    AsyncEngine,
    AsyncSession,
    async_sessionmaker,
    create_async_engine,
)

if TYPE_CHECKING:
    from twitter_api.config import Settings


_engine: AsyncEngine | None = None
_sessionmaker: async_sessionmaker[AsyncSession] | None = None


def create_engine_and_sessionmaker(
    settings: Settings,
) -> tuple[AsyncEngine, async_sessionmaker[AsyncSession]]:
    """エンジンとセッションファクトリを作って差し替える."""
    global _engine, _sessionmaker  # noqa: PLW0603 (シングルトン的に保持するため)

    _engine = create_async_engine(
        settings.database_url,
        echo=settings.database_echo,
        future=True,
        pool_pre_ping=True,
    )
    _sessionmaker = async_sessionmaker(
        bind=_engine,
        expire_on_commit=False,
        class_=AsyncSession,
    )
    return _engine, _sessionmaker


def get_engine() -> AsyncEngine:
    """初期化済みエンジンを取得."""
    if _engine is None:
        msg = "engine is not initialized; call create_engine_and_sessionmaker first"
        raise RuntimeError(msg)
    return _engine


def get_sessionmaker() -> async_sessionmaker[AsyncSession]:
    """初期化済みセッションファクトリを取得."""
    if _sessionmaker is None:
        msg = "sessionmaker is not initialized; call create_engine_and_sessionmaker first"
        raise RuntimeError(msg)
    return _sessionmaker
