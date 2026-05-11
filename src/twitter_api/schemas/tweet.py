"""Tweet の Pydantic スキーマ."""

from __future__ import annotations

from datetime import datetime

from pydantic import BaseModel, ConfigDict, Field


class CreateTweetRequest(BaseModel):
    """ツイート作成リクエスト."""

    model_config = ConfigDict(extra="forbid")

    author_id: str = Field(min_length=1)
    body: str = Field(min_length=1, max_length=140)


class CreateTweetResponse(BaseModel):
    """ツイート作成レスポンス."""

    id: str


class TweetDetailResponse(BaseModel):
    """ツイート詳細レスポンス."""

    tweet_id: str
    body: str
    author_id: str
    author_screen_name: str
    author_display_name: str
    created_at: datetime
    updated_at: datetime
