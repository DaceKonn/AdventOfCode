package helpers

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
