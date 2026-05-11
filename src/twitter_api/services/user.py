"""User サービス層.

Web API の慣習に従い、Router から呼び出される薄い service として実装する.
- 入力 (Pydantic schemas) を受け取り
- ORM モデルを操作
- ドメイン上のルールを満たさなければ AppError を送出
"""

from __future__ import annotations

from typing import TYPE_CHECKING

from sqlalchemy import select
from sqlalchemy.exc import IntegrityError

from twitter_api.exceptions import ConflictError, NotFoundError
from twitter_api.models import User

if TYPE_CHECKING:
    from sqlalchemy.ext.asyncio import AsyncSession

    from twitter_api.schemas import CreateUserRequest


class UserService:
    """User に関するビジネスロジック."""

    def __init__(self, session: AsyncSession) -> None:
        self._session = session

    async def create(self, payload: CreateUserRequest) -> User:
        """ユーザーを作成して返す.

        screen_name 重複時は ConflictError を送出.
        """
        user = User(
            screen_name=payload.screen_name,
            user_name=payload.user_name,
            bio=payload.bio,
        )
        self._session.add(user)
        try:
            await self._session.flush()
        except IntegrityError as e:
            await self._session.rollback()
            raise ConflictError("the screen name specified is already used") from e
        return user

    async def find_by_screen_name(self, screen_name: str) -> User:
        """screen_name でユーザーを取得.

        存在しなければ NotFoundError を送出.
        """
        stmt = select(User).where(User.screen_name == screen_name)
        result = await self._session.execute(stmt)
        user = result.scalar_one_or_none()
        if user is None:
            raise NotFoundError("user not found")
        return user
