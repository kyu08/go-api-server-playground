package domain

type StringVO[T any] string

func NewStringVO[T ~string](v string) T {
	return T(v)
}

func NewStringVOFromString[T ~string](s string) T {
	return T(s)
}

func (u StringVO[T]) Value() string {
	return string(u)
}
