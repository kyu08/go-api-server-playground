"""`python -m twitter_api` / `twitter-api` で起動するためのランチャ."""

from __future__ import annotations

import uvicorn

from twitter_api.config import get_settings


def main() -> None:
    """Uvicorn サーバを起動する."""
    settings = get_settings()
    uvicorn.run(
        "twitter_api.main:app",
        host=settings.host,
        port=settings.port,
        reload=False,
        log_config=None,
    )


if __name__ == "__main__":
    main()
