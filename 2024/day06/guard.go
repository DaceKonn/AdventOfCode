package main

import "github.com/DaceKonn/AdventOfCode/2024/day06/helpers"

const (
	facingUnknown = iota
	facingUp
	facingRight
	facingDown
	facingLeft
)

type guard struct {
	o      helpers.Object
	facing int
}

func newEmptyGuard() helpers.Object {
	return &guard{
		o:      helpers.NewEmptyObject(),
		facing: facingUnknown,
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

func (g *guard) SetSymbol(symbol rune) {
	g.o.SetSymbol(symbol)
}

func (g *guard) Copy() helpers.Object {
	return &guard{
		o:      g.o.Copy(),
		facing: g.GetFacing(),
	}
}
