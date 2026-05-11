"""Tweet サービス層."""

from __future__ import annotations

from typing import TYPE_CHECKING

from sqlalchemy import select
from sqlalchemy.orm import joinedload

from twitter_api.exceptions import NotFoundError
from twitter_api.models import Tweet, User

if TYPE_CHECKING:
    from sqlalchemy.ext.asyncio import AsyncSession

    from twitter_api.schemas import CreateTweetRequest


class TweetService:
    """Tweet に関するビジネスロジック."""

    def __init__(self, session: AsyncSession) -> None:
        self._session = session

    async def create(self, payload: CreateTweetRequest) -> Tweet:
        """ツイートを作成して返す.

        author が存在しなければ NotFoundError を送出.
        """
        author = await self._session.get(User, payload.author_id)
        if author is None:
            raise NotFoundError("user not found")

        tweet = Tweet(author_id=author.id, body=payload.body)
        self._session.add(tweet)
        await self._session.flush()
        return tweet

    async def get(self, tweet_id: str) -> Tweet:
        """ツイートを ID で取得 (author を eager load)."""
        stmt = select(Tweet).options(joinedload(Tweet.author)).where(Tweet.id == tweet_id)
        result = await self._session.execute(stmt)
        tweet = result.scalar_one_or_none()
        if tweet is None:
            raise NotFoundError("tweet not found")
        return tweet
