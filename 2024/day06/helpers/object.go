package helpers

const ObjectEmptySymbol rune = '?'

var ObjectEmptyId Id = NewIntId(-1)
var ObjectDefaultPoint Point = NewDefaultPoint(-1, -1)

type Object interface {
	GetOrigin() Point
	GetCurrent() Point
	GetSymbol() rune
	GetId() Id
	Copy() Object
	SetSymbol(symbol rune)
}

type DefaultObject struct {
	origin  Point
	current Point
	symbol  rune
	id      Id
}

func NewEmptyObject() Object {
	return &DefaultObject{
		origin:  ObjectDefaultPoint,
		current: ObjectDefaultPoint,
		symbol:  ObjectEmptySymbol,
		id:      ObjectEmptyId,
	}
}

func NewObject(origin, current Point, symbol rune, id Id) Object {
	return &DefaultObject{
		origin:  origin,
		current: current,
		symbol:  symbol,
		id:      id,
	}
}

func (do *DefaultObject) GetOrigin() Point {
	return do.origin
}

func (do *DefaultObject) GetCurrent() Point {
	return do.current
}

func (do *DefaultObject) GetSymbol() rune {
	return do.symbol
}

func (do *DefaultObject) GetId() Id {
	return do.id
}

func (do *DefaultObject) SetSymbol(symbol rune) {
	do.symbol = symbol
}

func (do *DefaultObject) Copy() Object {
	return &DefaultObject{
		origin:  do.GetOrigin().Copy(),
		current: do.GetCurrent().Copy(),
		id:      do.GetId(),
		symbol:  do.GetSymbol(),
	}
}
