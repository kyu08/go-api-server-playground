package id

import "github.com/google/uuid"

type ID string

func New() ID {
	return ID(uuid.New().String())
}

func (u ID) String() string {
	return string(u)
}
