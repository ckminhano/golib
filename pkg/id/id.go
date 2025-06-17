package id

import (
	"errors"

	"github.com/google/uuid"
)

// Id is a wrapper around uuid.UUID to provide a custom type for IDs.
// It is used to represent unique identifiers in the system.
// The zero value of Id is a valid ID, representing a nil UUID.
type Id uuid.UUID

// NewId creates a new Id with a random UUID.
func NewId() *Id {
	id := Id(uuid.New())
	return &id
}

// ToString converts the Id to a string representation of the UUID.
func (id *Id) ToString() string {
	return uuid.UUID(*id).String()
}

// ToUUID converts the Id to a uuid.UUID.
func (id *Id) ToUUID() uuid.UUID {
	return uuid.UUID(*id)
}

/*
FromString converts a string representation of a UUID to an Id.
It returns an error if the string is empty or s is a nil UUID in the form 0000000-0000-0000-0000-000000000000.
*/
func FromString(s string) (*Id, error) {
	if s == "" || s == uuid.Nil.String() {
		return nil, errors.New("string s cannot be empty")
	}
	id := Id(uuid.MustParse(s))
	return &id, nil
}
