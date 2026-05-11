"""User エンドポイントのテスト."""

from __future__ import annotations

from typing import TYPE_CHECKING

import pytest

if TYPE_CHECKING:
    from httpx import AsyncClient


UUID_LENGTH = 36


async def _create_user(
    client: AsyncClient,
    *,
    screen_name: str = "test_user",
    user_name: str = "Test User",
    bio: str = "This is a bio",
) -> dict[str, str]:
    res = await client.post(
        "/users",
        json={"screen_name": screen_name, "user_name": user_name, "bio": bio},
    )
    res.raise_for_status()
    body: dict[str, str] = res.json()
    return body


class TestCreateUser:
    """POST /users."""

    async def test_正常にユーザーを作成できる(self, client: AsyncClient) -> None:
        screen_name = "test_user"
        user_name = "Test User"
        bio = "This is a bio"

        res = await client.post(
            "/users",
            json={"screen_name": screen_name, "user_name": user_name, "bio": bio},
        )
        assert res.status_code == 201
        body = res.json()
        assert len(body["id"]) == UUID_LENGTH

        # 取得して確認
        get_res = await client.get(f"/users/{screen_name}")
        assert get_res.status_code == 200
        got = get_res.json()
        assert got["screen_name"] == screen_name
        assert got["user_name"] == user_name
        assert got["bio"] == bio

    @pytest.mark.parametrize(
        ("payload", "field"),
        [
            ({"screen_name": "", "user_name": "Test User", "bio": "bio"}, "screen_name"),
            ({"screen_name": "s", "user_name": "", "bio": "bio"}, "user_name"),
            ({"screen_name": "s", "user_name": "Test User", "bio": ""}, "bio"),
            ({"screen_name": "a" * 21, "user_name": "Test User", "bio": "bio"}, "screen_name"),
        ],
    )
    async def test_バリデーション違反は422(
        self,
        client: AsyncClient,
        payload: dict[str, str],
        field: str,
    ) -> None:
        res = await client.post("/users", json=payload)
        assert res.status_code == 422
        detail = res.json()["detail"]
        assert any(field in err["loc"] for err in detail)

    async def test_同じscreen_nameは409(self, client: AsyncClient) -> None:
        await _create_user(client, screen_name="dup_user")
        res = await client.post(
            "/users",
            json={"screen_name": "dup_user", "user_name": "Second User", "bio": "bio"},
        )
        assert res.status_code == 409
        assert res.json()["error"] == "conflict"


class TestFindUserByScreenName:
    """GET /users/{screen_name}."""

    async def test_存在するユーザーを取得できる(self, client: AsyncClient) -> None:
        created = await _create_user(client, screen_name="findme")
        res = await client.get("/users/findme")
        assert res.status_code == 200
        body = res.json()
        assert body["id"] == created["id"]
        assert body["screen_name"] == "findme"

    async def test_存在しないユーザーは404(self, client: AsyncClient) -> None:
        res = await client.get("/users/nonexistent_user")
        assert res.status_code == 404
        assert res.json()["error"] == "not_found"
