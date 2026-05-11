"""アプリケーション固有の例外と FastAPI 用ハンドラ.

Python の慣習として、業務エラーは独自例外として送出し、
FastAPI の例外ハンドラで HTTP ステータスに変換する.
"""

from __future__ import annotations

from typing import TYPE_CHECKING

import structlog
from fastapi import status
from fastapi.responses import JSONResponse

if TYPE_CHECKING:
    from fastapi import FastAPI, Request


log = structlog.get_logger()


class AppError(Exception):
    """アプリケーション例外の基底."""

    status_code: int = status.HTTP_500_INTERNAL_SERVER_ERROR
    error_code: str = "internal_error"

    def __init__(self, message: str = "internal server error") -> None:
        super().__init__(message)
        self.message = message


class NotFoundError(AppError):
    """リソースが見つからない."""

    status_code = status.HTTP_404_NOT_FOUND
    error_code = "not_found"


class ValidationError(AppError):
    """ドメインバリデーション違反.

    Pydantic の RequestValidationError とは別に、サービス層で発生する
    業務ロジックレベルのバリデーションエラーを表す.
    """

    status_code = status.HTTP_400_BAD_REQUEST
    error_code = "bad_request"


class ConflictError(AppError):
    """一意制約違反などの競合."""

    status_code = status.HTTP_409_CONFLICT
    error_code = "conflict"


def _app_error_handler(_: Request, exc: AppError) -> JSONResponse:
    if exc.status_code >= status.HTTP_500_INTERNAL_SERVER_ERROR:
        log.exception("unhandled application error", error_code=exc.error_code)
    else:
        log.warning("application error", error_code=exc.error_code, message=exc.message)
    return JSONResponse(
        status_code=exc.status_code,
        content={"error": exc.error_code, "message": exc.message},
    )


def _unhandled_exception_handler(_: Request, exc: Exception) -> JSONResponse:
    log.exception("unhandled exception", error=str(exc))
    return JSONResponse(
        status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
        content={"error": "internal_error", "message": "internal server error"},
    )


def register_exception_handlers(app: FastAPI) -> None:
    """FastAPI アプリに例外ハンドラを登録する."""
    app.add_exception_handler(AppError, _app_error_handler)  # type: ignore[arg-type]
    app.add_exception_handler(Exception, _unhandled_exception_handler)
