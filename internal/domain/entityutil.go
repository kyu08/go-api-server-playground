package domain

import (
	"github.com/google/uuid"
	"github.com/kyu08/go-api-server-playground/internal/apperrors"
)

// ID はPhantom Typeを使うことで型安全なIDを実現するための構造体。
// type ID stringのように定義してしまうとUserIDとTweetIDのように別のエンティティのIDを型で区別できない。
// （その結果引数の順番を間違えるなど間違った使われ方をしうる）
//
// また、type UserID string, type TweetID stringのように定義すると、
// tweetID := UserID(tweet.ID)のようにキャストできてしまうため安全性が不十分。
// それらの欠点を補うためにPhantom typeを利用して型安全性を高めている。
type ID[T any] struct {
	value uuid.UUID
	// _ T として定義してしまうと利用側(e.g. User構造体の定義場所)で
	// invalid recursive type User (see details) [InvalidDeclCycle]
	// のようなエラーになるためポインタ型にしている。
	_ *T
}

func NewID[T any]() ID[T] {
	//nolint:exhaustruct // Phantom typeなので致し方なし
	return ID[T]{value: uuid.New()}
}

func NewFromString[T any](s string) (ID[T], error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return ID[T]{}, apperrors.WithStack(err)
	}
	//nolint:exhaustruct // Phantom typeなので全フィールド指定できないのは仕方ない。
	return ID[T]{value: u}, nil
}

func (u ID[T]) String() string {
	return u.value.String()
}
