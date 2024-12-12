package helpers

import (
	"fmt"
)

type Id struct {
	indId      int
	stringId   string
	isStringId bool
}

func NewStringId(id string) Id {
	return Id{-1, id, true}
}

func NewIntId(id int) Id {
	return Id{id, "", false}
}

func (id Id) String() string {
	switch id.isStringId {
	case true:
		return id.stringId
	case false:
		return fmt.Sprint(id.indId)
	}
	panic("I can't handle to string for Id!")
}
