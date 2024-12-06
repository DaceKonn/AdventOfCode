package helpers

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

const ObjectEmptySymbol rune = '?'

var ObjectEmptyId Id = NewIntId(-1)
var ObjectDefaultPoint Point = NewDefaultPoint(-1, -1)

type Object interface {
	GetOrigin() Point
	GetCurrent() Point
	GetSymbol() rune
	GetId() Id

	SetSymbol(symbol rune)
	SetOrigin(origin Point)
	SetCurrent(current Point)
	SetId(id Id)

	Copy() Object
}

type DefaultObject struct {
	origin  Point
	current Point
	symbol  rune
	id      Id
}

// SetCurrent implements Object.
func (do *DefaultObject) SetCurrent(current Point) {
	do.current = current
}

// SetId implements Object.
func (do *DefaultObject) SetId(id Id) {
	do.id = id
}

// SetOrigin implements Object.
func (do *DefaultObject) SetOrigin(origin Point) {
	do.origin = origin
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

func LogObjectDebug(o Object) {
	log.Debug().
		Type("type", o).
		Str("rune", string(o.GetSymbol())).
		Str("origin", fmt.Sprint(o.GetOrigin())).
		Str("current", fmt.Sprint(o.GetCurrent())).
		Str("id", fmt.Sprint(o.GetId())).
		Msg("Object debug")
}
