"""ヘルスチェック router."""

from __future__ import annotations

from fastapi import APIRouter

router = APIRouter(tags=["health"])


@router.get("/healthz", summary="ヘルスチェック")
async def healthz() -> dict[str, str]:
    """生死確認用エンドポイント."""
    return {"message": "twitter"}
