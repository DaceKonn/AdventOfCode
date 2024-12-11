package main

import (
	"github.com/DaceKonn/AdventOfCode/2024/day06/helpers"
	"github.com/rs/zerolog/log"
)

const (
	facingUnknown = iota
	facingUp
	facingRight
	facingDown
	facingLeft
)

func facingToString(facing int) string {
	switch facing {
	case facingUnknown:
		return "unknown"
	case facingUp:
		return "up"
	case facingRight:
		return "right"
	case facingDown:
		return "down"
	case facingLeft:
		return "left"
	default:
		return "undefined"
	}
}

type guard struct {
	o         helpers.Object
	facing    int
	hasExited bool
}

// GetFlags implements helpers.Object.
func (g *guard) GetFlags() map[string]bool {
	return g.o.GetFlags()
}

// SetFlag implements helpers.Object.
func (g *guard) SetFlag(key string, value bool) {
	g.o.SetFlag(key, value)
}

// SetCurrent implements helpers.Object.
func (g *guard) SetCurrent(current helpers.Point) {
	g.o.SetCurrent(current)
}

// SetId implements helpers.Object.
func (g *guard) SetId(id helpers.Id) {
	g.o.SetId(id)
}

// SetOrigin implements helpers.Object.
func (g *guard) SetOrigin(origin helpers.Point) {
	g.o.SetOrigin(origin)
}

func newEmptyGuard() *guard {
	return &guard{
		o:         helpers.NewEmptyObject(),
		facing:    facingUnknown,
		hasExited: false,
	}
}

func (g *guard) GetOrigin() helpers.Point {
	return g.o.GetOrigin()
}

func (g *guard) GetCurrent() helpers.Point {
	return g.o.GetCurrent()
}

func (g *guard) GetId() helpers.Id {
	return g.o.GetId()
}

func (g *guard) GetSymbol() rune {
	return g.o.GetSymbol()
}

func (g *guard) GetFacing() int {
	return g.facing
}

func (g *guard) HasExited() bool {
	return g.hasExited
}

func (g *guard) SetExited(v bool) {
	g.hasExited = v
}

func (g *guard) SetSymbol(symbol rune) {
	g.o.SetSymbol(symbol)
	g.facing = symbolToFacing(symbol)
	log.Debug().
		Str("facing", facingToString(g.facing)).
		Str("id", g.o.GetId().String()).
		Msg("Guard facing updated")
}

func (g *guard) Copy() helpers.Object {
	return &guard{
		o:      g.o.Copy(),
		facing: g.GetFacing(),
	}
}

func (g *guard) SetFacing(facing int) {
	g.facing = facing
	g.o.SetSymbol(facingToSymbol(facing))
	log.Debug().
		Str("facing", facingToString(facing)).
		Str("id", g.o.GetId().String()).
		Msg("Guard facing updated")
}

func symbolToFacing(symbol rune) int {
	switch symbol {
	case '^':
		return facingUp
	case '>':
		return facingRight
	case 'V', 'v':
		return facingDown
	case '<':
		return facingLeft
	default:
		return facingUnknown

	}
}

func facingToSymbol(facing int) rune {
	switch facing {
	case facingUnknown:
		return '?'
	case facingUp:
		return '^'
	case facingRight:
		return '>'
	case facingDown:
		return 'v'
	case facingLeft:
		return '<'
	default:
		return '!'
	}
}
