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
	GetFlags() map[string]bool

	SetSymbol(symbol rune)
	SetOrigin(origin Point)
	SetCurrent(current Point)
	SetId(id Id)
	SetFlag(key string, value bool)
}

type DefaultObject struct {
	origin  Point
	current Point
	symbol  rune
	id      Id
	flags   map[string]bool
}

// GetFlags implements Object.
func (do *DefaultObject) GetFlags() map[string]bool {
	return do.flags
}

// SetFlags implements Object.
func (do *DefaultObject) SetFlag(key string, value bool) {
	do.flags[key] = value
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
		flags:   make(map[string]bool),
	}
}

func NewObject(origin, current Point, symbol rune, id Id) Object {
	return &DefaultObject{
		origin:  origin,
		current: current,
		symbol:  symbol,
		id:      id,
		flags:   make(map[string]bool),
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

func LogObjectDebug(o Object) {
	log.Debug().
		Type("type", o).
		Str("rune", string(o.GetSymbol())).
		Str("origin", fmt.Sprint(o.GetOrigin())).
		Str("current", fmt.Sprint(o.GetCurrent())).
		Str("id", fmt.Sprint(o.GetId())).
		Str("flags", fmt.Sprint(o.GetFlags())).
		Msg("Object debug")
}

func LogObjectInfo(o Object) {
	log.Info().
		Type("type", o).
		Str("rune", string(o.GetSymbol())).
		Str("origin", fmt.Sprint(o.GetOrigin())).
		Str("current", fmt.Sprint(o.GetCurrent())).
		Str("id", fmt.Sprint(o.GetId())).
		Str("flags", fmt.Sprint(o.GetFlags())).
		Msg("Object debug")
}
