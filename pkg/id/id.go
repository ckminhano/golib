package id

import (
	"errors"

	"github.com/google/uuid"
)

type Id uuid.UUID

func NewId() *Id {
	return &Id{}
}

func (id *Id) ToString() string {
	return uuid.UUID(*id).String()
}

func (id *Id) ToUUID() uuid.UUID {
	return uuid.UUID(*id)
}

func FromString(s string) (*Id, error) {
	if s == "" {
		return nil, errors.New("string s cannot be empty")
	}
	id := Id(uuid.MustParse(s))
	return &id, nil
}
