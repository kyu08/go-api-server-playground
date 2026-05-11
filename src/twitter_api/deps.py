"""FastAPI 用の Dependency Injection ヘルパ.

Web API の慣習として、DB セッションや Service のインスタンス化は ``Depends`` を介して
リクエストスコープで提供する.
"""

from __future__ import annotations

from collections.abc import AsyncIterator
from typing import Annotated

from fastapi import Depends
from sqlalchemy.ext.asyncio import AsyncSession

from twitter_api.db import get_sessionmaker
from twitter_api.services import TweetService, UserService


async def get_session() -> AsyncIterator[AsyncSession]:
    """リクエストごとの AsyncSession を払い出す.

    レスポンス成功時は commit、例外時は rollback. router 側は明示的に commit しない.
    """
    sessionmaker = get_sessionmaker()
    async with sessionmaker() as session:
        try:
            yield session
        except Exception:
            await session.rollback()
            raise
        else:
            await session.commit()


SessionDep = Annotated[AsyncSession, Depends(get_session)]


def get_user_service(session: SessionDep) -> UserService:
    """UserService をリクエストスコープで生成."""
    return UserService(session)


def get_tweet_service(session: SessionDep) -> TweetService:
    """TweetService をリクエストスコープで生成."""
    return TweetService(session)


UserServiceDep = Annotated[UserService, Depends(get_user_service)]
TweetServiceDep = Annotated[TweetService, Depends(get_tweet_service)]
