"""User ORM モデル."""

from __future__ import annotations

import uuid
from datetime import UTC, datetime
from typing import TYPE_CHECKING

from sqlalchemy import DateTime, String
from sqlalchemy.orm import Mapped, mapped_column, relationship

from twitter_api.db.base import Base

if TYPE_CHECKING:
    from twitter_api.models.tweet import Tweet


def _now_utc() -> datetime:
    return datetime.now(UTC)


class User(Base):
    """ユーザー."""

    __tablename__ = "users"

    id: Mapped[str] = mapped_column(
        String(36),
        primary_key=True,
        default=lambda: str(uuid.uuid4()),
    )
    screen_name: Mapped[str] = mapped_column(String(20), unique=True, index=True, nullable=False)
    user_name: Mapped[str] = mapped_column(String(20), nullable=False)
    bio: Mapped[str] = mapped_column(String(160), nullable=False)
    created_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True),
        default=_now_utc,
        nullable=False,
    )
    updated_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True),
        default=_now_utc,
        onupdate=_now_utc,
        nullable=False,
    )

    tweets: Mapped[list[Tweet]] = relationship(
        back_populates="author",
        cascade="all, delete-orphan",
        lazy="raise",
    )
