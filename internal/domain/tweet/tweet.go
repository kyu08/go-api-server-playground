package tweet

import (
	"time"

	"github.com/kyu08/go-api-server-playground/internal/domain"
	"github.com/kyu08/go-api-server-playground/internal/domain/user"
)

type Tweet struct {
	ID        domain.ID[Tweet]
	AuthorID  domain.ID[user.User]
	body      Body // ツイート本文
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Tweet) Body() Body {
	return t.body
}

// NewTweet はツイート作成時に使用するコンストラクタ。
func NewTweet(authorID domain.ID[user.User], body string) (*Tweet, error) {
	now := time.Now()
	return newTweet(domain.NewID[Tweet](), authorID, body, now, now)
}

// NewFromDTO は主にDTOからエンティティを生成する際に使用されることを想定し、IDを外から受け取るようにしている。
func NewFromDTO(idString, authorIDString, body string, createdAt, updatedAt time.Time) (*Tweet, error) {
	id, err := domain.NewFromString[Tweet](idString)
	if err != nil {
		return nil, err
	}

	authorID, err := domain.NewFromString[user.User](authorIDString)
	if err != nil {
		return nil, err
	}

	return newTweet(id, authorID, body, createdAt, updatedAt)
}

func newTweet(id domain.ID[Tweet], authorID domain.ID[user.User], body string, createdAt, updatedAt time.Time) (*Tweet, error) {
	b, err := NewBody(body)
	if err != nil {
		return nil, err
	}

	return &Tweet{
		ID:        id,
		AuthorID:  authorID,
		body:      b,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
