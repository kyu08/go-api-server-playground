"""SQLAlchemy 2.0 Declarative Base."""

from __future__ import annotations

from sqlalchemy.orm import DeclarativeBase


class Base(DeclarativeBase):
    """全 ORM モデルの基底クラス."""
