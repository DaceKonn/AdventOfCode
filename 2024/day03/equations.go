package main

import (
	"fmt"
	"strings"
)

const (
	mulMaxNumLength int = 3
)

func mulCommand() string {
	return "mul"
}

func mulOpeners() string {
	return "("
}

func mulClosures() string {
	return ")"
}

type calculator interface {
	calculate() int
	print() string
}

type element struct {
	Value int
}

func newElement(value int) element {
	// fmt.Printf("\t\tcreating element: %d\n", value)
	return element{Value: value}
}

func (e element) calculate() int {
	return e.Value
}

func (e element) print() string {
	return fmt.Sprintf("%d", e.Value)
}

type add struct {
	elements []calculator
}

func newAdd() add {
	// fmt.Println("\t\tcreating add")
	return add{}
}

func (a *add) addElement(el calculator) {
	// fmt.Println("\t\t\t adding element")
	a.elements = append(a.elements, el)
}

func (a add) calculate() int {
	var base = 0
	for _, el := range a.elements {
		base += el.calculate()
	}
	return base
}

func (a add) print() string {
	var outSlice []string
	for _, el := range a.elements {
		outSlice = append(outSlice, el.print())
	}
	switch len(outSlice) {
	case 0:
		return "emptyAdd"
	case 1:
		return fmt.Sprintf("singleAdd(%s)", outSlice[0])
	default:
		return strings.Join(outSlice, " + ")
	}
}

type multiply struct {
	elements []calculator
}

func newMultiply() multiply {
	// fmt.Println("\t\tcreating multiply")
	return multiply{}
}

func (m *multiply) addElement(el calculator) {
	// fmt.Println("\t\t\t adding element")
	m.elements = append(m.elements, el)
}

func (m multiply) calculate() int {
	var base int = 1
	for _, el := range m.elements {
		base *= el.calculate()
	}
	return base
}

func (m multiply) print() string {
	var outSlice []string
	for _, el := range m.elements {
		outSlice = append(outSlice, el.print())
	}
	switch len(outSlice) {
	case 0:
		return "emptyMultiply"
	case 1:
		return fmt.Sprintf("singleMultiply(%s)", outSlice[0])
	default:
		return strings.Join(outSlice, " * ")
	}
}
