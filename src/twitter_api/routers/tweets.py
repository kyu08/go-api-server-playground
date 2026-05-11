"""Tweet router."""

from __future__ import annotations

from fastapi import APIRouter, status

from twitter_api.deps import TweetServiceDep
from twitter_api.schemas import CreateTweetRequest, CreateTweetResponse, TweetDetailResponse

router = APIRouter(prefix="/tweets", tags=["tweets"])


@router.post(
    "",
    response_model=CreateTweetResponse,
    status_code=status.HTTP_201_CREATED,
    summary="ツイート作成",
)
async def create_tweet(
    payload: CreateTweetRequest,
    service: TweetServiceDep,
) -> CreateTweetResponse:
    """新しいツイートを作成する."""
    tweet = await service.create(payload)
    return CreateTweetResponse(id=tweet.id)


@router.get(
    "/{tweet_id}",
    response_model=TweetDetailResponse,
    summary="ツイート詳細取得",
)
async def get_tweet(
    tweet_id: str,
    service: TweetServiceDep,
) -> TweetDetailResponse:
    """ツイート詳細を取得する."""
    tweet = await service.get(tweet_id)
    return TweetDetailResponse(
        tweet_id=tweet.id,
        body=tweet.body,
        author_id=tweet.author_id,
        author_screen_name=tweet.author.screen_name,
        author_display_name=tweet.author.user_name,
        created_at=tweet.created_at,
        updated_at=tweet.updated_at,
    )
