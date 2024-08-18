package id

import "github.com/google/uuid"

type ID string

func New() ID {
	return ID(uuid.New().String())
}

// TODO: add UT
func (u ID) Validate() error {
	panic("unimplemented")
}

func (u ID) String() string {
	return string(u)
}
