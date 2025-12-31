package domain

type StringVO[T any] struct {
	value string
	// _ T として定義してしまうと利用側(e.g. User構造体の定義場所)で
	// invalid recursive type User (see details) [InvalidDeclCycle]
	// のようなエラーになるためポインタ型にしている。
	_ *T
}

func NewStringVO[T any](v string) StringVO[T] {
	//nolint:exhaustruct // Phantom typeなので致し方なし
	return StringVO[T]{value: v}
}

func NewStringVOFromString[T any](s string) StringVO[T] {
	//nolint:exhaustruct // Phantom typeなので全フィールド指定できないのは仕方ない。
	return StringVO[T]{value: s}
}

func (u StringVO[T]) Value() string {
	return u.value
}
