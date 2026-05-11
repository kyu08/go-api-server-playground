"""Tweet エンドポイントのテスト."""

from __future__ import annotations

import uuid
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from httpx import AsyncClient


UUID_LENGTH = 36


async def _create_user(client: AsyncClient, *, screen_name: str | None = None) -> dict[str, str]:
    res = await client.post(
        "/users",
        json={
            "screen_name": screen_name or str(uuid.uuid4())[:20],
            "user_name": "Test User",
            "bio": "bio",
        },
    )
    res.raise_for_status()
    body: dict[str, str] = res.json()
    return body


class TestCreateTweet:
    """POST /tweets."""

    async def test_140文字マルチバイトでも作成できる(self, client: AsyncClient) -> None:
        user = await _create_user(client)
        body = "あ" * 140
        res = await client.post("/tweets", json={"author_id": user["id"], "body": body})
        assert res.status_code == 201
        tweet_id = res.json()["id"]
        assert len(tweet_id) == UUID_LENGTH

        get_res = await client.get(f"/tweets/{tweet_id}")
        assert get_res.status_code == 200
        got = get_res.json()
        assert got["tweet_id"] == tweet_id
        assert got["body"] == body

    async def test_140文字ASCIIでも作成できる(self, client: AsyncClient) -> None:
        user = await _create_user(client)
        body = "a" * 140
        res = await client.post("/tweets", json={"author_id": user["id"], "body": body})
        assert res.status_code == 201

    async def test_author_idが空だと422(self, client: AsyncClient) -> None:
        res = await client.post("/tweets", json={"author_id": "", "body": "hi"})
        assert res.status_code == 422

    async def test_bodyが空だと422(self, client: AsyncClient) -> None:
        user = await _create_user(client)
        res = await client.post("/tweets", json={"author_id": user["id"], "body": ""})
        assert res.status_code == 422

    async def test_bodyが長すぎると422(self, client: AsyncClient) -> None:
        user = await _create_user(client)
        res = await client.post("/tweets", json={"author_id": user["id"], "body": "a" * 141})
        assert res.status_code == 422

    async def test_存在しないユーザーで作成すると404(self, client: AsyncClient) -> None:
        res = await client.post(
            "/tweets",
            json={"author_id": str(uuid.uuid4()), "body": "hi"},
        )
        assert res.status_code == 404
        assert res.json()["error"] == "not_found"


class TestGetTweet:
    """GET /tweets/{tweet_id}."""

    async def test_詳細を取得できる(self, client: AsyncClient) -> None:
        user = await _create_user(client, screen_name="napper_pro")
        # ユーザーを取り直してフルレスポンスを得る
        await client.get("/users/napper_pro")

        post = await client.post(
            "/tweets",
            json={"author_id": user["id"], "body": "おひるねなう"},
        )
        tweet_id = post.json()["id"]

        res = await client.get(f"/tweets/{tweet_id}")
        assert res.status_code == 200
        body = res.json()
        assert body["tweet_id"] == tweet_id
        assert body["body"] == "おひるねなう"
        assert body["author_id"] == user["id"]
        assert body["author_screen_name"] == "napper_pro"
        assert body["author_display_name"] == "Test User"
        assert "created_at" in body
        assert "updated_at" in body

    async def test_存在しないtweet_idは404(self, client: AsyncClient) -> None:
        res = await client.get(f"/tweets/{uuid.uuid4()}")
        assert res.status_code == 404
        assert res.json()["error"] == "not_found"
