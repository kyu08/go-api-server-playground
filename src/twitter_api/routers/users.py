"""User router."""

from __future__ import annotations

from fastapi import APIRouter, status

from twitter_api.deps import UserServiceDep
from twitter_api.schemas import CreateUserRequest, CreateUserResponse, UserResponse

router = APIRouter(prefix="/users", tags=["users"])


@router.post(
    "",
    response_model=CreateUserResponse,
    status_code=status.HTTP_201_CREATED,
    summary="ユーザー作成",
)
async def create_user(
    payload: CreateUserRequest,
    service: UserServiceDep,
) -> CreateUserResponse:
    """新しいユーザーを作成する."""
    user = await service.create(payload)
    return CreateUserResponse(id=user.id)


@router.get(
    "/{screen_name}",
    response_model=UserResponse,
    summary="screen_name でユーザー取得",
)
async def find_user_by_screen_name(
    screen_name: str,
    service: UserServiceDep,
) -> UserResponse:
    """screen_name でユーザーを 1 件取得する."""
    user = await service.find_by_screen_name(screen_name)
    return UserResponse.model_validate(user)
