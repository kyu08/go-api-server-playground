"""User の Pydantic スキーマ.

Python Web API の慣習として、入力バリデーションは Pydantic に寄せる.
"""

from __future__ import annotations

from pydantic import BaseModel, ConfigDict, Field


class CreateUserRequest(BaseModel):
    """ユーザー作成リクエスト."""

    model_config = ConfigDict(extra="forbid")

    screen_name: str = Field(min_length=1, max_length=20)
    user_name: str = Field(min_length=1, max_length=20)
    bio: str = Field(min_length=1, max_length=160)


class CreateUserResponse(BaseModel):
    """ユーザー作成レスポンス."""

    id: str


class UserResponse(BaseModel):
    """ユーザー取得レスポンス."""

    model_config = ConfigDict(from_attributes=True)

    id: str
    screen_name: str
    user_name: str
    bio: str
