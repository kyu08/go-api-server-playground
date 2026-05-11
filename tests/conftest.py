"""共通の pytest フィクスチャ.

Web API のテストの慣習として、以下のフィクスチャを提供する:
- アプリ全体の lifespan を切り替えてインメモリ SQLite を使う ``app``
- ``httpx.AsyncClient`` ベースの ``client``
"""

from __future__ import annotations

from collections.abc import AsyncIterator
from contextlib import asynccontextmanager

import pytest
from fastapi import FastAPI
from httpx import ASGITransport, AsyncClient
from sqlalchemy.ext.asyncio import (
    AsyncSession,
    async_sessionmaker,
    create_async_engine,
)

from twitter_api.config import Settings
from twitter_api.db import session as db_session
from twitter_api.db.base import Base
from twitter_api.main import create_app


@pytest.fixture
def settings() -> Settings:
    """テスト用設定 (in-memory SQLite, JSON ログ無効)."""
    return Settings(
        database_url="sqlite+aiosqlite:///:memory:",
        database_echo=False,
        log_json=False,
    )


@pytest.fixture
async def app(settings: Settings) -> AsyncIterator[FastAPI]:
    """テストごとに独立した in-memory DB を持つ FastAPI アプリ.

    各テストで使い切りの sessionmaker を差し込み、テーブルを毎回生成する.
    """
    # NOTE: in-memory SQLite はコネクションごとに別 DB になるため、StaticPool で 1 本に固定する.
    from sqlalchemy.pool import StaticPool

    engine = create_async_engine(
        settings.database_url,
        echo=False,
        connect_args={"check_same_thread": False},
        poolclass=StaticPool,
    )
    sessionmaker = async_sessionmaker(bind=engine, expire_on_commit=False, class_=AsyncSession)

    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)

    # 通常の create_engine_and_sessionmaker をバイパスして直接差し込む
    db_session._engine = engine
    db_session._sessionmaker = sessionmaker

    @asynccontextmanager
    async def test_lifespan(_: FastAPI) -> AsyncIterator[None]:
        yield

    fastapi_app = create_app(settings=settings)
    fastapi_app.router.lifespan_context = test_lifespan
    try:
        yield fastapi_app
    finally:
        await engine.dispose()
        db_session._engine = None
        db_session._sessionmaker = None


@pytest.fixture
async def client(app: FastAPI) -> AsyncIterator[AsyncClient]:
    """テスト用 HTTPX クライアント (ASGI in-process)."""
    transport = ASGITransport(app=app)
    async with AsyncClient(transport=transport, base_url="http://test") as ac:
        yield ac
