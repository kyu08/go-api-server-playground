"""FastAPI アプリケーションのエントリポイント.

慣習的な構成として、``create_app`` でアプリを組み立て、Uvicorn から呼び出す.
テストでもこの関数を再利用する.
"""

from __future__ import annotations

from collections.abc import AsyncIterator
from contextlib import asynccontextmanager
from typing import TYPE_CHECKING

import structlog
from fastapi import FastAPI

from twitter_api import __version__
from twitter_api.config import get_settings
from twitter_api.db import create_engine_and_sessionmaker
from twitter_api.db.base import Base
from twitter_api.exceptions import register_exception_handlers
from twitter_api.logging import configure_logging
from twitter_api.routers import health, tweets, users

if TYPE_CHECKING:
    from twitter_api.config import Settings


log = structlog.get_logger()


@asynccontextmanager
async def _lifespan(app: FastAPI) -> AsyncIterator[None]:
    """アプリのライフサイクルで DB エンジンを管理.

    開発時には毎回テーブルを作成する. Production では Alembic で migrate する想定.
    """
    settings: Settings = app.state.settings
    engine, _ = create_engine_and_sessionmaker(settings)

    # 開発用 SQLite ではこのまま使えるよう、起動時にテーブルを生成しておく.
    # ここで models を import するのは、Base.metadata に登録するため.
    from twitter_api import models  # noqa: F401, PLC0415

    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)

    log.info("application started", version=__version__)
    try:
        yield
    finally:
        await engine.dispose()
        log.info("application stopped")


def create_app(settings: Settings | None = None) -> FastAPI:
    """FastAPI インスタンスを生成して返す."""
    settings = settings or get_settings()
    configure_logging(settings)

    app = FastAPI(
        title="Twitter API Playground",
        version=__version__,
        lifespan=_lifespan,
    )
    app.state.settings = settings

    register_exception_handlers(app)

    app.include_router(health.router)
    app.include_router(users.router)
    app.include_router(tweets.router)

    return app


app = create_app()
